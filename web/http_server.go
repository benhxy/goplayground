package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strings"
)

func main() {
	http.HandleFunc("/", helloHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if strings.HasSuffix(path, "@benhu.dev") {
		fmt.Fprintf(w, "Hello Ben's friend %s \n", strings.TrimSuffix(path, "@benhu.dev")[1:])
	} else {
		fmt.Fprintf(w, "Hello user %v \n", path[1:])
	}
}
