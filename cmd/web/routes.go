package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rhc07/basic-web-app/pkg/config"
	"github.com/rhc07/basic-web-app/pkg/handlers"
)

func Routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	// mux.Get("/euros", http.HandlerFunc(handlers.Repo.Euros))

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/euros", handlers.Repo.Euros)

	return mux
}
