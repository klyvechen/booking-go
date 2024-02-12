package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github/klyvechen/booking-go/pkg/config"
	"github/klyvechen/booking-go/pkg/handlers"
	"github/klyvechen/booking-go/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":4000"

var app config.AppConfig

var session *scs.SessionManager

func main() {

	// change this tho true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
