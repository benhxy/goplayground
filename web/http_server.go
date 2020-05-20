package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(rw http.ResponseWriter, r *http.Request) {
	response := []byte("hello world!")
	rw.Write(response)
}
