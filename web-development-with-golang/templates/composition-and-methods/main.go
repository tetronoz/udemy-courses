package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

func (p person) DoubleAge() int {
	return 2 * p.Age
}

func (p person) TakesArgs(x int) int {
	return 2 * x
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	p1 := person{
		Name: "James Bond",
		Age:  42,
	}

	if err := tpl.Execute(os.Stdout, p1); err != nil {
		log.Fatalln(err)
	}
}