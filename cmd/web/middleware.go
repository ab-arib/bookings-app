package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf is a midddleware to generate CSFRToken handler
func NoSurf(next http.Handler) http.Handler {
	csfrHandler := nosurf.New(next)

	csfrHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",              //implement to all path
		Secure:   app.InProduction, //non https
		SameSite: http.SameSiteLaxMode,
	})

	return csfrHandler
}

// SessionLoad load and save session every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
