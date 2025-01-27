package controllers

import (
	"net/http"

	"cerescms.hapaxredux.net/internal/views"
	"github.com/go-chi/chi/v5"
)

var (
	tmplMgr = views.NewTmplEngOnDemand("../../ui/templates")
)

func webRoutes(r chi.Router) {
	r.Get("/", index)
	r.Route("/users", func(r chi.Router) {
		r.Get("/signup", userSignup)
		r.Get("/login", userLogin)
		r.Get("/logout", userLogout)
	})
	r.Get("/dashboard", pageDashboard)
}

func index(w http.ResponseWriter, r *http.Request) {
	data := views.NewTmplHydr(r)
	tmplMgr.Render(w, http.StatusOK, "home.tmpl.html", data)
}

func userSignup(w http.ResponseWriter, r *http.Request) {
	data := views.NewTmplHydr(r)
	tmplMgr.Render(w, http.StatusOK, "signup.tmpl.html", data)
}

func userLogin(w http.ResponseWriter, r *http.Request) {
	data := views.NewTmplHydr(r)
	tmplMgr.Render(w, http.StatusOK, "signup.tmpl.html", data)
}

func userLogout(w http.ResponseWriter, r *http.Request) {
	data := views.NewTmplHydr(r)
	tmplMgr.Render(w, http.StatusOK, "signup.tmpl.html", data)
}

func pageDashboard(w http.ResponseWriter, r *http.Request) {
	data := views.NewTmplHydr(r)
	tmplMgr.Render(w, http.StatusOK, "dashboard.tmpl.html", data)
}
