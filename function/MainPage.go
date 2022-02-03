package groupie

import (
	"html/template"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../static/index.html"))

	data := APIRequestArtist()

	tmpl.Execute(w, data)
}
