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

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	mux.Get("/", app.Home)
	mux.Get("/health", app.Health)
	mux.Get("/test-patterns", app.TestPatterns)

	mux.Group(func(gr chi.Router) {
		gr.Route("/api", func(r chi.Router) {
			r.Get("/dog-from-factory", app.NewDogFromFactory)
			r.Get("/cat-from-factory", app.NewCatFromFactory)

			r.Get("/dog-from-abstract-factory", app.NewDogFromAbstractFactory)
			r.Get("/cat-from-abstract-factory", app.NewCatFromAbstractFactory)
		})
	})

	mux.Get("/{page}", app.ShowPage)

	return mux
}
