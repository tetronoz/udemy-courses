package main

import (
	"io"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Index Page")
}

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Dog Page")

}

func me(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Sergey!")
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenAndServe failed %v", err)
	}
}