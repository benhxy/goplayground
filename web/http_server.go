package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

func main() {
	http.HandleFunc("/", helloHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	regex := regexp.MustCompile("^(.+)@golang.org$")

	path := r.URL.Path
	match := regex.Find([]byte(path))
	if match != nil {
		fmt.Fprintf(w, "Hello gopher %v \n", string(match[1:]))
	} else {
		fmt.Fprintf(w, "Hello user %v \n", path)
	}
}
