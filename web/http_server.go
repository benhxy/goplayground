package web

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
	regex := regexp.MustCompile("^(.+)@benhu.dev$")
	path := r.URL.Path
	match := regex.FindStringSubmatch(path)

	if match != nil && len(match) > 1 {
		fmt.Fprintf(w, "Hello Ben's friend %v \n", match[1][1:])
	} else {
		fmt.Fprintf(w, "Hello user %v \n", path[1:])
	}
}
