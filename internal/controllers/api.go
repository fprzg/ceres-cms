package controllers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"cerescms.hapaxredux.net/internal/services"
	"github.com/go-chi/chi/v5"
)

var basePath = "/home/ubu24-t480/Code/ceres-cms/examples"

func apiRoutes(r chi.Router) {
	r.Get("/res/{projID}/*", readResource)
	r.Post("/res/{projID}/*", updateResource)
	r.Get("/render/{projID}/{*}", renderResource)
}

func getResourcePath(proj, res string) string {
	return filepath.Join(basePath, proj, "i18n", res)
}

func readResource(w http.ResponseWriter, r *http.Request) {
	projID := chi.URLParam(r, "projID")
	resID := chi.URLParam(r, "*")

	path := getResourcePath(projID, resID)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic(err)
	}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buff, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var ser services.JsonSer
	asJson, err := ser.FromYaml(buff)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(asJson)
}

func updateResource(w http.ResponseWriter, r *http.Request) {
	projID := chi.URLParam(r, "projID")
	resID := chi.URLParam(r, "*")

	path := getResourcePath(projID, resID)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var ser services.JsonSer
	asJson, err := ser.FromRequest(r)
	if err != nil {
		panic(err)
	}

	_, err = io.Writer.Write(file, asJson)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}

func renderResource(w http.ResponseWriter, r *http.Request) {

}
