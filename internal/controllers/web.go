package controllers

import (
	"net/http"

	"cerescms.hapaxredux.net/internal/views"
	"github.com/go-chi/chi/v5"
)

var (
	tmplMgr = views.NewTemplateMgr()
)

func webRoutes(r chi.Router) {
	r.Get("/", pageIndex)
	r.Get("/signup", pageSignup)
	r.Get("/login", pageLogin)
	r.Get("/logout", pageLogout)
}

func pageIndex(w http.ResponseWriter, r *http.Request) {
	data := views.NewTemplateHydr(r)
	tmplMgr.RenderView(w, http.StatusOK, "home.tmpl.html", data)
}

func pageSignup(w http.ResponseWriter, r *http.Request) {
	data := views.NewTemplateHydr(r)
	tmplMgr.RenderView(w, http.StatusOK, "signup.tmpl.html", data)
}

func pageLogin(w http.ResponseWriter, r *http.Request) {
	data := views.NewTemplateHydr(r)
	tmplMgr.RenderView(w, http.StatusOK, "signup.tmpl.html", data)
}

func pageLogout(w http.ResponseWriter, r *http.Request) {
	data := views.NewTemplateHydr(r)
	tmplMgr.RenderView(w, http.StatusOK, "signup.tmpl.html", data)
}
