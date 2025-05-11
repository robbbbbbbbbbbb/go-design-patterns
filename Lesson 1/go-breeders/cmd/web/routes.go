package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(60 * time.Second))

	fileServer := http.FileServer(http.Dir("../../static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	mux.Get("/", app.Home)
	mux.Get("/health", app.Health)
	mux.Get("/{page}", app.ShowPage)
	mux.Get("/new-dog", app.NewDogFromFactory)
	mux.Get("/new-cat", app.NewCatFromFactory)
	return mux
}
