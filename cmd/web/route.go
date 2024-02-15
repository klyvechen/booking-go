package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github/klyvechen/booking-go/pkg/config"
	"github/klyvechen/booking-go/pkg/handlers"
	"net/http"
)

func routes(config *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// middle ware
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/general-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)
	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJson)
	mux.Get("/contact", handlers.Repo.Contact)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
