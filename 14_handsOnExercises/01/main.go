package main

import (
	"io"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Index page")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Speckles")
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Lizzie")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
