package views

import (
	"bytes"
	"fmt"
	"io/fs"
	"net/http"
	"path/filepath"
	"text/template"
	"time"

	"cerescms.hapaxredux.net/ui"
)

func readableDate() int {
	return 666
}

type TemplateMgr struct {
	cache     map[string]*template.Template
	functions template.FuncMap
}

func NewTemplateMgr() TemplateMgr {
	mgr := TemplateMgr{
		cache: map[string]*template.Template{},
		functions: template.FuncMap{
			"readableDate": readableDate,
		},
	}

	pages, err := fs.Glob(ui.Files, "templates/pages/*.tmpl.html")
	if err != nil {
		panic(err)
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"templates/base.tmpl.html",
			"templates/partials/nav.tmpl.html",
			page,
		}

		ts, err := template.New(name).Funcs(mgr.functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			panic(err)
		}

		mgr.cache[name] = ts
	}

	return mgr
}

type TemplateHydr struct {
	CurrentYear     int
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

func NewTemplateHydr(r *http.Request) *TemplateHydr {
	td := &TemplateHydr{
		CurrentYear:     time.Now().Year(),
		IsAuthenticated: false,
		CSRFToken:       "none",
	}

	if r != nil {
		td.Flash = ""
		td.IsAuthenticated = false //app.isAuthenticated(r)
	}

	return td
}

func (mgr *TemplateMgr) HydrTemplate(page string, data *TemplateHydr) *bytes.Buffer {
	ts, ok := mgr.cache[page]
	if !ok {
		panic(fmt.Errorf("template '%s' does not exist", page))
	}

	buff := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buff, "base", data)
	if err != nil {
		panic(err)
	}

	return buff
}

func (mgr *TemplateMgr) RenderView(w http.ResponseWriter, status int, page string, data *TemplateHydr) {
	buff := mgr.HydrTemplate(page, data)

	w.WriteHeader(status)
	buff.WriteTo(w)
}
