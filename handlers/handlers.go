package handlers

import (
	"net/http"

	"github.com/torenware/sawler-course/config"
	"github.com/torenware/sawler-course/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewHandlers(a *config.AppConfig) {
	Repo = &Repository{App: a}
}

func (s *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.pages.tmpl")
}

func (s *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.pages.tmpl")
}
