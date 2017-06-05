package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type familyMember struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	myFamilyMember := familyMember{
		Name:  "Adam",
		Motto: "When's dinner?",
	}
	err := tpl.Execute(os.Stdout, myFamilyMember)
	if err != nil {
		log.Fatalln(err)
	}
}
