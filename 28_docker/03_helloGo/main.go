package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	log.Fatalln(http.ListenAndServe(":80", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, this is GO inside a docker container!")
}
