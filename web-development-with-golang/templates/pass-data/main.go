package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/tpl.gohtml"))
}

type data struct {
	Age int
	Name string
}

func main() {
	d := data{Age: 42, Name: "Sergey"}

	if err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", d); err != nil {
		log.Fatalln(err)
	}

	fmt.Println()

	tpl, err := tpl.ParseFiles("templates/slice.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	sages := []string{"Ghandi", "MLK"}

	if err := tpl.ExecuteTemplate(os.Stdout, "slice.gohtml", sages); err != nil {
		log.Fatalln(err)
	}

	fmt.Println()

	tpl, err = tpl.ParseFiles("templates/map.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	m := map[string]string{
		"India": "Ghandi",
		"America": "MLK",
	}

	if err := tpl.ExecuteTemplate(os.Stdout, "map.gohtml", m); err != nil {
		log.Fatalln(err)
	}

}