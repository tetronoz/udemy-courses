package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

var fm = template.FuncMap {
	"fdateDMY": dayMonthYear,
	"fdateMDY": monthDayYear,
}

func dayMonthYear(t time.Time) string {
	return t.Format("02-01-2006")
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}


func main() {
	if err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now()); err != nil {
		log.Fatalln(err)
	}
}