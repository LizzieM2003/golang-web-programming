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
	Admin bool
}

func main() {
	a := familyMember{"Adam", "What's for dinner?", false}
	b := familyMember{"", "I'm invisible", true}
	c := familyMember{"Darren", "Tea dear?", true}

	members := []familyMember{a, b, c}

	err := tpl.Execute(os.Stdout, members)

	if err != nil {
		log.Fatalln(err)
	}

}
