package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/test/", TestAPI)
	http.ListenAndServe(":8085", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func TestAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a test API, response from , %s", r.URL.Path[1:])
}
