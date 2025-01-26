package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	staticDir := "./"
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, "en")
	})
	http.HandleFunc("/es", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, "es")
	})
	http.HandleFunc("/fr", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, "fr")
	})
	http.HandleFunc("/jp", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, "jp")
	})

	port := ":4500"
	log.Printf("Starting server on %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("ListenAndServe failed:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request, lang string) {
	tmpl, err := template.ParseFiles("index.tmpl.html")
	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		log.Println("Error loading template:", err)
		return
	}

	path := fmt.Sprintf("i18n/%s/index.yaml", lang)
	hydr, err := loadYaml(path)
	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		log.Println("Error loading template:", err)
		return
	}

	err = tmpl.Execute(w, hydr)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
	}
}

func loadYaml(filename string) (*map[string]interface{}, error) {
	var contents map[string]interface{}
	asYaml, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(asYaml, &contents)
	if err != nil {
		return nil, err
	}

	return &contents, nil
}
