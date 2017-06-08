package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/apply", apply)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func about(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "about.gohtml", nil)
}

func contact(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "contact.gohtml", nil)
}

func apply(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
	} else {
		tpl.ExecuteTemplate(w, "apply.gohtml", nil)
	}
}
