package groupie

import (
	"fmt"
	"html/template"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	tmpl := template.Must(template.ParseFiles("../static/index.html"))

=======
	tmpl := template.Must(template.ParseFiles("../static/index.html"))	
>>>>>>> e2197bf6abc7936a9f0e46f5c319c453e9c49ef5
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
