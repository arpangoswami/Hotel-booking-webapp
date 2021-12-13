package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/arpangoswami/Hotel-booking-webapp/pkg/config"
	"github.com/arpangoswami/Hotel-booking-webapp/pkg/handlers"
	"github.com/arpangoswami/Hotel-booking-webapp/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig
var sessionManager *scs.SessionManager

func main() {

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can't create a template cache")
	}
	//change this when in production
	app.InProduction = false

	//initialising session manager
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Name = "session_id"
	sessionManager.IdleTimeout = 20 * time.Minute
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.Secure = app.InProduction
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode

	app.Session = sessionManager
	newRepo := handlers.NewRepository(&app)
	handlers.NewHandlers(newRepo)

	app.TemplateCache = templateCache
	app.UseCache = false

	render.NewTemplates(&app)
	fmt.Println(fmt.Printf("Listening on port %s", portNumber))
	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = server.ListenAndServe()

	if err != nil {
		log.Fatal("Error Listen and serve: ", err.Error())
	}
}
