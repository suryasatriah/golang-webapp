package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/suryasatriah/learn-go/pkg/config"
	"github.com/suryasatriah/learn-go/pkg/handlers"
	"github.com/suryasatriah/learn-go/pkg/render"
)

const portNumber = ":3000"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	//Change this to false when in production
	app.InProduction = true

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Secure = app.InProduction
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Persist = true

	app.Session = session

	//create template cache
	tc, err := render.CreateTemplateChache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	var repo = handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s", portNumber)

	serve := http.Server{
		Addr:    portNumber,
		Handler: route(&app),
	}

	err = serve.ListenAndServe()
	log.Fatal(err)
}
