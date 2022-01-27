package main

import (
	"net/http"
	g  "groupie/function"
)
func main() {	
	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe("localhost:8080", nil)
	g.MainHandler(http ReponseWriter, r *http.Request)
}