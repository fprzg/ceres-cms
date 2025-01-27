package controllers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/", webRoutes)

	r.Handle("/static/*", dirNoListing(http.FileServer(http.Dir("../../ui/"))))

	r.Route("/v1", func(r chi.Router) {
		r.Route("/api", apiRoutes)
	})

	return r
}

func dirNoListing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Stat(filepath.Join("../../ui/", r.URL.Path))
		if err == nil && file.IsDir() {
			http.NotFound(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
