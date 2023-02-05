package cmd

import (
	"net/http"
	"text/template"
)

func LoginCreate(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("site2/login.html")

	if err != nil {
		panic(err)
	}

	template.ExecuteTemplate(w, "login", nil)
}
