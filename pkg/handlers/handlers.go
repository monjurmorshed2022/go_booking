package handlers

import (
	"net/http"

	"github.com/monjurmorshed2022/go_booking/pkg/config"
	"github.com/monjurmorshed2022/go_booking/pkg/models"
	"github.com/monjurmorshed2022/go_booking/pkg/render"
)

//Repo the repository used by the handler
var Repo *Repository

//Repository is the Repository Type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// SetHandlers sets the Repository for the Handler
func SetHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(rw http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(rw, "home.page.html", &models.TemplateDate{})
}

//About is the about page handler
func (m *Repository) About(rw http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Passed From Handler............"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(rw, "about.page.html", &models.TemplateDate{
		StringMap: stringMap,
	})
}
