package controllers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/", webRoutes)
	r.Route("/v1", func(r chi.Router) {
		r.Route("/api", apiRoutes)
	})

	return r
}
