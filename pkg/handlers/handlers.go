package handlers

import (
	"net/http"

	"github.com/rhc07/basic-web-app/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home_page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about_page.html")
}

func Euros(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "euros_page.html")
}
