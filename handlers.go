package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home_page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about_page.html")
}

func Euros(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "euros_page.html")
}
