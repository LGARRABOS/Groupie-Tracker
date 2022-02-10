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
		m := r.PostForm
		st := m["ID"]
		if len(st) != 0 {
			number, apifile := SeparateString(st[0])
			data1 := SubApiVerif(apifile, number)
			fmt.Println(data1)
		}

	}
	tmpl.Execute(w, data)
}