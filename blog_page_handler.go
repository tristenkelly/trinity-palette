package main

import "net/http"

func (cfg *apiConfig) blogHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "blog.html")
}
