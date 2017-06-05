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

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Family    []familyMember
	Transport []car
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

	car1 := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	car2 := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	cars := []car{car1, car2}

	data := items{
		Family:    myFamily,
		Transport: cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
