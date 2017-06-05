package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type familyMember struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func main() {
	a := familyMember{
		Name:  "Adam",
		Motto: "When's dinner?",
	}

	b := familyMember{
		Name:  "Hannah",
		Motto: "I have too much homework.",
	}

	c := familyMember{
		Name:  "Speckles",
		Motto: "Woof woof!",
	}

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

	data := struct {
		Family    []familyMember
		Transport []car
	}{
		[]familyMember{a, b, c},
		[]car{car1, car2},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
