package main

import (
	"net/http"

	"github.com/arpangoswami/Hotel-booking-webapp/pkg/config"
	"github.com/arpangoswami/Hotel-booking-webapp/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurfCSRF)
	mux.Use(LoadSession)
	mux.Get("/", handlers.Repository.Home)
	mux.Get("/about", handlers.Repository.About)
	return mux
}
