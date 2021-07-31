package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)


func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func main() {
	mux := httprouter.New()
	mux.GET("/", index)
}