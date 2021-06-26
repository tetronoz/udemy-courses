package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)


var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree (s string) string {
	s = strings.TrimSpace(s)
	return s[:3]
}

type sage struct {
	Name string
	Moto string
}

func main() {
	b := sage{
		Name: "Buddha",
		Moto: "The belief of no belief",
	}

	g := sage{
		Name: "Ghandi",
		Moto: "Be the change",
	}

	sages := []sage{b, g}

	data := struct {
		Wisdom    []sage
	}{
		sages,
	}

	if err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data); err != nil {
		log.Fatalln(err)
	}

}