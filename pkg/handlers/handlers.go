package handlers

import (
	"net/http"

	"github.com/arpangoswami/Hotel-booking-webapp/pkg/config"
	"github.com/arpangoswami/Hotel-booking-webapp/pkg/models"
	"github.com/arpangoswami/Hotel-booking-webapp/pkg/render"
)

type RepositoryPattern struct {
	App *config.AppConfig
}

var Repository *RepositoryPattern

//NewRepository initialises Repository
func NewRepository(appLocal *config.AppConfig) *RepositoryPattern {
	newRepo := RepositoryPattern{
		App: appLocal,
	}
	return &newRepo
}

//NewHandlers sets the repository for the handlers
func NewHandlers(repoLocal *RepositoryPattern) {
	Repository = repoLocal
}

//Home is the home page handler
func (repoLocal *RepositoryPattern) Home(rw http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	repoLocal.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(rw, "home.page.html", &models.TemplateData{})
}

//About is the about route handler
func (repoLocal *RepositoryPattern) About(rw http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["headerSubText"] = "Random About Text"
	remoteIP := repoLocal.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["ip_address"] = remoteIP

	render.RenderTemplate(rw, "about.page.html", &models.TemplateData{StringData: stringMap})
}
