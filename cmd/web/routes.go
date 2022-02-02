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
	mux.Get("/rooms/generals-quarters", handlers.Repository.GeneralsQuarters)
	mux.Get("/rooms/admirals-abode", handlers.Repository.AdmiralsAbode)
	mux.Get("/rooms/colonels-batcave", handlers.Repository.ColonelsBatcave)
	mux.Get("/make-reservation", handlers.Repository.MakeReservation)
	mux.Get("/contact", handlers.Repository.Contact)
	mux.Get("/search-availability", handlers.Repository.SearchAvailability)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
