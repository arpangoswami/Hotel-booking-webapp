package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

//NoSurfCSRF adds CSRF protection to all POST requests
func NoSurfCSRF(next http.Handler) http.Handler {
	csrfToken := nosurf.New(next)
	csrfToken.SetBaseCookie(http.Cookie{
		Name:     "Cookies",
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   24 * 60 * 60,
	})
	return csrfToken
}

//LoadSession saves the session on every request
func LoadSession(next http.Handler) http.Handler {
	return sessionManager.LoadAndSave(next)
}
