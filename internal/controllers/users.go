package controllers

import (
	"net/http"

	"cerescms.net/internal/views"
)

func home(w http.ResponseWriter, r *http.Request) {
	data := views.NewTemplateHydr(r)
	tmplMgr.RenderView(w, http.StatusOK, "home.tmpl.html", data)
}
