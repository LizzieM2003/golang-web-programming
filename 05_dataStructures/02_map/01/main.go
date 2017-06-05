package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	myFamily := map[string]string{
		"husband":  "Dazzie",
		"son":      "Adam",
		"daughter": "Hannah",
		"dog":      "Speckles"}
	err := tpl.Execute(os.Stdout, myFamily)

	if err != nil {
		log.Fatalln(err)
	}
}
