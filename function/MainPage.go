package groupie

import (
	"net/http"
	"html/template"
)

func MainHandler(http ReponseWriter, r *http.Request ) {
	tmpl := template.Must(template.ParseFiles("../static/index.html"))
	
	APIRequest()
	data := ArtisStruct {
		Tab: ArtistTab
	}
	tmpl.Execute(w, data)
}
