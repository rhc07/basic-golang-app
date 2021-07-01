package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rhc07/basic-web-app/pkg/config"
	"github.com/rhc07/basic-web-app/pkg/handlers"
	"github.com/rhc07/basic-web-app/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/euros", handlers.Euros)

	fmt.Println(fmt.Sprintf("Starting on port number: %s", portNumber))
	http.ListenAndServe(portNumber, nil)

}
