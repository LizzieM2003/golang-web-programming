package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{
	"fdateDMY": dayMonthYear,
}

func dayMonthYear(t time.Time) string {
	return t.Format("02/01/2006")
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())

	if err != nil {
		log.Fatalln(err)
	}
}
