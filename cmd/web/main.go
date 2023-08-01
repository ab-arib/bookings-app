package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	//importing the package from another directory
	"github.com/ab-arib/bookings-app/pkg/config"
	"github.com/ab-arib/bookings-app/pkg/handlers"
	"github.com/ab-arib/bookings-app/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":3001"

var app config.AppConfig // implement global app config
var session *scs.SessionManager

// main is the main application function
func main() {
	app.InProduction = false // set the production state

	// initiate web session
	session = scs.New()
	session.Lifetime = 24 * time.Hour // how long session last
	session.Cookie.Persist = true     // set the data to not be clear after the user close the page
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // non https

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false // use cache used as test case, if in dev mode no longer read from template cache

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println("Starting application server on port", portNumber)
	// serve http listen service
	srv := &http.Server{
		Addr:    portNumber,
		Handler: Routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
