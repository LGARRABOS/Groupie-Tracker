package groupie

import (
	"net/http"
	"html/template"
	"fmt"
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
		// m := r.PostForm
		// id := m["ID"]

		// if (lens(ID) != 0) {
			
		// }
		
	}

	
	tmpl.Execute(w, data)
}
