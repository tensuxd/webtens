package cmd

import (
	"net/http"
	"text/template"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("site2/home.html")

	if err != nil {
		panic(err)
	}

	tmpl.ExecuteTemplate(w, "home", nil)
}
