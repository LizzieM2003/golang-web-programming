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

type familyMember struct {
	Name  string
	Motto string
}

func main() {
	Adam := familyMember{
		Name:  "Adam",
		Motto: "What's for dinner?",
	}
	Hannah := familyMember{
		Name:  "Hannah",
		Motto: "I have too much homework.",
	}
	Darren := familyMember{
		Name:  "Darren",
		Motto: "Tea dear?",
	}
	Speckles := familyMember{
		Name:  "Speckles",
		Motto: "Pigeons on roof!",
	}

	myFamily := []familyMember{Adam, Hannah, Darren, Speckles}

	err := tpl.Execute(os.Stdout, myFamily)
	if err != nil {
		log.Fatalln(err)
	}
}
