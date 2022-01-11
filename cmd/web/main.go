package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/monjurmorshed2022/go_booking/pkg/config"
	"github.com/monjurmorshed2022/go_booking/pkg/handlers"
	"github.com/monjurmorshed2022/go_booking/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8081"

var sessionManager *scs.SessionManager
var app config.AppConfig

//main is the main application function...
func main() {

	//change this to true when in production
	app.InProduction = false
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.InProduction

	app.Session = sessionManager

	templateCatch, err := render.CreateTemplateCatch()

	if err != nil {
		log.Fatal("Cannot Create Template Catch", err)
	}

	app.TemplateCache = templateCatch

	repository := handlers.NewRepo(&app)
	handlers.SetHandlers(repository)

	render.SetTemplate(&app)

	fmt.Printf("%s\n", fmt.Sprintf("Starting the application on port %s", portNumber))

	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	log.Fatal(err)
}
