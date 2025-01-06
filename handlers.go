package main

import (
	"net/http"
	"text/template"
)

//where you can select an account
//or create an account if you don't have any
func renderHomePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/home.html"))
	tmpl.Execute(w, nil)
}
