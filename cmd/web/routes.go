package main

import (
	"net/http"

	"github.com/ab-arib/bookings-app/pkg/config"
	"github.com/ab-arib/bookings-app/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// http handler routes
func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	//implement middleware
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/home", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/devide", handlers.Devide)

	return mux
}
