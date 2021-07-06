package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/rhc07/basic-web-app/pkg/config"
	"github.com/rhc07/basic-web-app/pkg/handlers"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/euros", http.HandlerFunc(handlers.Repo.Euros))

	return mux
}
