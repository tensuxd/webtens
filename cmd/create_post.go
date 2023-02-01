package cmd

import (
	"net/http"
	"text/template"
)

func Create(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("site2/create_post.html")

	if err != nil {
		panic(err)
	}

	tmpl.ExecuteTemplate(w, "create_post", nil)

}
