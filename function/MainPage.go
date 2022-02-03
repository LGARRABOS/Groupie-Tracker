package groupie

import (
	"fmt"
	"html/template"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../static/index.html"))

	data := APIRequestArtist()

	switch r.Method {
	case "GET":
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Printf("POST METHOD = %v\n", r.PostForm)

	}

	tmpl.Execute(w, data)
}
