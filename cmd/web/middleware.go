package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

//SessionLoad Loads and Saves the Session on Every  Request
func SessionLoad(next http.Handler) http.Handler {
	return sessionManager.LoadAndSave(next)
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}
