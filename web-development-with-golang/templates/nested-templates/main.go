package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	if err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", 42); err != nil {
		log.Fatalln(err)
	}
}