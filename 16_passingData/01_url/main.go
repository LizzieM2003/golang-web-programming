package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	x := req.FormValue("p")
	fmt.Fprintln(w, "Do my search:", v, x)
}

// visit this page:
// http://localhost:8080/?q=dog&p=cat
