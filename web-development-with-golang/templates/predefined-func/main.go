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

	xs := []string{"zero", "one", "two", "three", "four", "five"}

	if err := tpl.Execute(os.Stdout, xs); err != nil {
		log.Fatalln(err)
	}
}