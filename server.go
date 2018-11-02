package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	w.Write([]byte("hello world\n"))
}

func bucketHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"id": 1, "name": "my bucket 1"}`)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h2>Index page</h2>")
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
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler)
	r.HandleFunc("/api/buckets/{id}", bucketHandler)
	r.HandleFunc("/api/buckets", createBucketHandler).Methods("POST")
	r.HandleFunc("/", indexHandler)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
