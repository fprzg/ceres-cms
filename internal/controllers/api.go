package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"cerescms.hapaxredux.net/internal/services"
	"github.com/go-chi/chi/v5"
)

func apiRoutes(r chi.Router) {
	r.Get("/res/{projectID}/*", getResource)
}

func getResource(w http.ResponseWriter, r *http.Request) {
	//projectID := chi.URLParam(r, "projectID")
	resourceID := chi.URLParam(r, "*")

	basePath := "/home/ubu24-t480/Code/ceres-cms/examples/i18n/"
	fullPath := filepath.Join(basePath, resourceID)

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		panic(err)
	}

	file, err := os.Open(fullPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buff, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	yp := services.YamlParser{}

	asJson, err := yp.ToJSON(buff)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "%s", asJson)
}
