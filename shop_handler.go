package main

import (
	"net/http"
)

func shopHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "shop.html")

}
