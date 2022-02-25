package groupie

import (
	"fmt"
	"html/template"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../static/map.html"))
	js := template.Must(template.ParseFiles("../static/page.html"))
	data := APIRequestArtist()
	var start int
	

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
		choice := m["choix"]
		create := m["creation"]
		first := m["first"]
		loc := m["location"]
		data = Filter(loc, first, create, choice)
		redir := m ["redir"]
		if len(redir) != 0 {
			if redir[0] == "home" {
				start = 0
			} else if redir[0] == "page" {
				start = 1
			}
		}
		if len(st) != 0 {
			number, apifile := SeparateString(st[0])
			data1 := SubApiVerif(apifile, number)
			fmt.Println(data1)
		}
	}
	if start == 0 {
		tmpl.Execute(w, data)
	} else if start == 1 {
		js.Execute(w, data)
	}
}
