package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// 오직 테스트 목적을 위한 함수
func Sum(a ...int) int {
	sum := 0
	for _, i := range a {
		sum += i
	}

	return sum
}

func bucketHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"id": 1, "name": "my bucket 1"}`)
}

type Bucket struct {
	Name string `json:"name"`
}

func createBucketHandler(w http.ResponseWriter, r *http.Request) {
	var bucket Bucket
	err := json.NewDecoder(r.Body).Decode(&bucket)

	if err != nil {
		fmt.Println("Error decoding JSON = ", err)
		return
	}

	log.Println(r.Method, r.URL, bucket.Name)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bucket)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/buckets/{id}", bucketHandler)
	router.HandleFunc("/api/buckets", createBucketHandler).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
