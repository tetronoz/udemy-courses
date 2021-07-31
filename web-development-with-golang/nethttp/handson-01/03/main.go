package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func index(w http.ResponseWriter, r *http.Request) {
	page := "Index"
	tpl.ExecuteTemplate(w, "index.gohtml", page)
}

func dog(w http.ResponseWriter, r *http.Request) {
	page := "Dog"
	tpl.ExecuteTemplate(w, "index.gohtml", page)
}

func me(w http.ResponseWriter, r *http.Request) {
	name := "Sergey"
	tpl.ExecuteTemplate(w, "me.gohtml", name)
}

func main() {

	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenAndServe failed %v", err)
	}
}