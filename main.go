package main

import (
	"log"
	"net/http"

	"go.mod/cmd"
)

func main() {

	log.Println("start")

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", cmd.Handler)
	http.HandleFunc("/create_post", cmd.Create)
	http.HandleFunc("/showPost", cmd.ShowPost)
	http.HandleFunc("/create_login", cmd.LoginCreate)
	http.ListenAndServe(":3030", nil)
}
