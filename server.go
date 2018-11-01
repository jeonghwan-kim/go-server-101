package main

import (
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

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler)
	r.HandleFunc("/api/buckets/{id}", bucketHandler)
	r.HandleFunc("/", indexHandler)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
