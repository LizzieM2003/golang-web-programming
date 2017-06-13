package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	FName string
	LName string
	Items []string
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/marshal", mshl)
	http.HandleFunc("/encode", encd)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, req *http.Request) {
	s := `
	<!DOCTYPE html>
	<html>
	<body>
		<p>You are at foo</p>
	</body>
	</html>
	`
	w.Write([]byte(s))
}

func mshl(w http.ResponseWriter, req *http.Request) {
	p1 := person{"James", "Bond", []string{"Gun", "Suit", "British sense of humour"}}

	json, err := json.Marshal(p1)

	if err != nil {
		log.Println(err)
	}

	w.Write(json)
}

func encd(w http.ResponseWriter, req *http.Request) {
	p1 := person{"James", "Bond", []string{"Gun", "Suit", "British sense of humour"}}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}
