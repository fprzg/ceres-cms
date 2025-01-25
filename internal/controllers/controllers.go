package controllers

import (
	"net/http"

	"cerescms.net/internal/views"
	"github.com/julienschmidt/httprouter"
)

var (
	tmplMgr = views.NewTemplateMgr()
)

func Router() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", pageIndex)
	router.GET("/user/signup", pageSignup)
	router.GET("/user/login", pageLogin)
	router.GET("/user/logout", pageLogout)

	return router
}

func pageIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := views.NewTemplateHydr(r)
	tmplMgr.RenderView(w, http.StatusOK, "home.tmpl.html", data)
}

func pageSignup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := views.NewTemplateHydr(r)
	tmplMgr.RenderView(w, http.StatusOK, "signup.tmpl.html", data)
}

func pageLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := views.NewTemplateHydr(r)
	tmplMgr.RenderView(w, http.StatusOK, "signup.tmpl.html", data)
}

func pageLogout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := views.NewTemplateHydr(r)
	tmplMgr.RenderView(w, http.StatusOK, "signup.tmpl.html", data)
}
