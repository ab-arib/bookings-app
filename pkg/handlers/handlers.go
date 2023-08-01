package handlers

import (
	"fmt"
	"net/http"

	"github.com/ab-arib/bookings-app/pkg/calculates"
	"github.com/ab-arib/bookings-app/pkg/config"
	"github.com/ab-arib/bookings-app/pkg/models"
	"github.com/ab-arib/bookings-app/pkg/render"
)

/** repository patern used to swap component with minimal changes on the code base */

// repo the repository used by the handlers
var Repo *Repository

type Repository struct {
	App config.AppConfig
}

// creates new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: *a,
	}
}

// sets the rrepository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler -- repository set as receiver () so the function can used it's value
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "This is the home page") //Fprinf print response to the webpage - where Println print response to the console
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp) //put remote ip address in app config

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// some logic for test passing param value to template
	stringMap := make(map[string]string)
	stringMap["test"] = "this is testing"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip") //accessing the app config value
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Devide page handler
func Devide(w http.ResponseWriter, r *http.Request) {
	f, err := calculates.DevideValue(100.0, 50.0)

	// error handler
	if err != nil {
		fmt.Fprintf(w, "Cannot devide by 0")
		return
	}

	fmt.Fprintf(w, "%f devided by %f is %f", 100.0, 50.0, f) //"%f" is a place holder require to logs float in a string.
}
