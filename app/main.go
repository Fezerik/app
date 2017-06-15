package main

import (
	"github.com/Fezerik/app/controller"
	"net/http"
)

func main() {

	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/signup", controller.SignUp)
	http.HandleFunc("/activate", controller.Activate)
	http.HandleFunc("/login", controller.Login)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
