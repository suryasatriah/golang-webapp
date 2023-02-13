package handlers

import (
	"net/http"

	"github.com/suryasatriah/learn-go/pkg/config"
	"github.com/suryasatriah/learn-go/pkg/model"
	"github.com/suryasatriah/learn-go/pkg/render"
)

var Repo *Repository

//TemplateData store datatype that want to be sent by handlers to template

type Repository struct {
	App *config.AppConfig
}

// Create repository for the handler
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Set repository for the handler
func NewHandler(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.pages.tmpl", &model.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//Create logic
	Stringmap := make(map[string]string)
	Stringmap["1"] = "Hello again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	Stringmap["remote_ip"] = remoteIP

	//Render template
	render.RenderTemplate(w, "about.pages.tmpl", &model.TemplateData{
		StringMap: Stringmap,
	})
}
