package handlers

import (
	"net/http"

	"github.com/rhc07/basic-web-app/pkg/config"
	"github.com/rhc07/basic-web-app/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// This is the Repository Type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}

func (m *Repository) Euros(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "euros.page.html")
}
