package views

import (
	"bytes"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
	"time"
)

func readableDate() int {
	return 666
}

type TmplHydr struct {
	I18n            map[string]interface{}
	CurrentYear     int
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

func NewTmplHydr(r *http.Request) *TmplHydr {
	td := &TmplHydr{
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

type TmplEng interface {
	Render(w http.ResponseWriter, status int, page string, data *TmplHydr)
}

type TmplEngOnDemand struct {
	templateDir string
	functions   template.FuncMap
}

func templateFunctions() template.FuncMap {
	return template.FuncMap{
		"readableDate": readableDate,
	}
}

func NewTmplEngOnDemand(templateDir string) *TmplEngOnDemand {
	return &TmplEngOnDemand{
		templateDir: templateDir,
		functions:   templateFunctions(),
	}
}

func (e *TmplEngOnDemand) Render(w http.ResponseWriter, status int, page string, data *TmplHydr) error {
	patterns := []string{
		filepath.Join(e.templateDir, "base.tmpl.html"),
		filepath.Join(e.templateDir, "partials/nav.tmpl.html"),
		filepath.Join(e.templateDir, "pages", page),
	}

	ts, err := template.New(page).Funcs(e.functions).ParseFiles(patterns...)
	if err != nil {
		panic(err)
	}

	var buff bytes.Buffer
	if err := ts.ExecuteTemplate(&buff, "base", data); err != nil {
		return err
	}

	w.WriteHeader(status)
	buff.WriteTo(w)

	return nil
}

type TmplEngCache struct {
	cache     map[string]*template.Template
	functions template.FuncMap
}

func NewTmplEngCache(templateDir string) TmplEngCache {
	e := TmplEngCache{
		cache:     map[string]*template.Template{},
		functions: templateFunctions(),
	}

	pages, err := filepath.Glob(filepath.Join(templateDir, "pages/*.tmpl.html"))
	if err != nil {
		panic(err)
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			filepath.Join(templateDir, "base.tmpl.html"),
			filepath.Join(templateDir, "partials/nav.tmpl.html"),
			page,
		}

		ts, err := template.New(name).Funcs(e.functions).ParseFiles(patterns...)
		if err != nil {
			panic(err)
		}

		e.cache[name] = ts
	}

	return e
}

func (e *TmplEngCache) Render(w http.ResponseWriter, status int, page string, data *TmplHydr) error {
	ts, err := e.cache[page]
	if !err {
		return fmt.Errorf("template not found")
	}

	var buff bytes.Buffer
	if err := ts.ExecuteTemplate(&buff, "base", data); err != nil {
		return err
	}

	w.WriteHeader(status)
	buff.WriteTo(w)

	return nil
}
