package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server is running on port 8000")
	ConnectDatabase()
	route := mux.NewRouter()

	log.Fatal(http.ListenAndServe(":8000", route))
}
