package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/foo", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	s := `Foo says "Hi!"`
	w.Write([]byte(s))
}
