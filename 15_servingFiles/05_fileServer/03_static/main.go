package main

import (
	"log"
	"net/http"
)

// Need to have an index.html file so that your static website directories and files are not exposed when
// go to root route

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
