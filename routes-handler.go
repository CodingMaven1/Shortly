package main

import "net/http"

// RenderHome route to render the home page
func RenderHome(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./views/index.html")
}
