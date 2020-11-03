package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//AddRoutes function for adding routes...
func AddRoutes(router *mux.Router) {

	log.Println("Loading routes")

	fs := http.FileServer(http.Dir("./public/"))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	router.HandleFunc("/", RenderHome)

	router.HandleFunc("/getshorturl", MakeShortURL).Methods("POST")

	router.HandleFunc("/{shortcode}", RedirectURL).Methods("GET")

	log.Println("Routes loaded")
}
