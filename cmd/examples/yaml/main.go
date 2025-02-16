package main

import (
	"encoding/json"
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

func main() {
	yamlData := `
pages:
  Home: "Inicio"
  About Us: "Sobre Nosotros"
hero:
  title: "Deja que tu hogar sea único"
  paragraph: "Hay muchas variaciones de los pasajes de lorem Ipsum disponibles, variaciones de los pasajes."
  button: "Obtén una estimación gratuita"
steps:
  - icon: "ri-compasses-2-line"
    title: "Planificación del Proyecto"
    paragraph: "Hay muchas variaciones de los pasajes de lorem Ipsum disponibles, variaciones de los pasajes."
  - icon: "ri-magic-line"
    title: "Obtención de Materiales"
    paragraph: "Hay muchas variaciones de los pasajes de lorem Ipsum disponibles, variaciones de los pasajes."
  - icon: "ri-tools-line"
    title: "Obtención de Materiales"
    paragraph: "Hay muchas variaciones de los pasajes de lorem Ipsum disponibles, variaciones de los pasajes."
about:
  text_title: "Creamos el arte de vivir con estilo de manera estilosa"
  text_paragraph: "Es un hecho establecido que un lector se distraerá con el contenido legible de una página al mirar su diseño, los puntos de uso que tienen más o menos normalidad."
  phone_number: "(+1) 884 234 2343"
  phone_paragraph: "Llámanos en cualquier momento"
testimonials:
  - img: "/static/img/testimonial/01.png"
    name: "Nattasha Mith"
    location: "Greenville, USA"
    paragraph: "Lorem Ipsum es simplemente texto de relleno de la industria tipográfica. Ipsum lo ha sido."
  - img: "/static/img/testimonial/02.png"
    name: "Nattasha Mith"
    location: "Greenville, USA"
    paragraph: "Lorem Ipsum es simplemente texto de relleno de la industria tipográfica. Ipsum lo ha sido."
  - img: "/static/img/testimonial/03.png"
    name: "Nattasha Mith"
    location: "Greenville, USA"
    paragraph: "Lorem Ipsum es simplemente texto de relleno de la industria tipográfica. Ipsum lo ha sido."
brands:
  - "/static/img/brands/01.svg"
  - "/static/img/brands/02.svg"
  - "/static/img/brands/03.svg"
  - "/static/img/brands/04.svg"
  - "/static/img/brands/05.svg"
projects:
  text:
    title: "Sigue Nuestros Proyectos"
    paragraph: "Es un hecho establecido que un lector se distraerá con el contenido legible de una página al mirar su diseño."
  list:
    - name: "Cocina Moderna"
      desc: "Decoración/Arquitectura"
      img: "/static/img/work/01.png"
    - name: "Cocina Moderna"
      desc: "Decoración/Arquitectura"
      img: "/static/img/work/02.png"
    - name: "Cocina Moderna"
      desc: "Decoración/Arquitectura"
      img: "/static/img/work/03.png"
    - name: "Cocina Moderna"
      desc: "Decoración/Arquitectura"
      img: "/static/img/work/04.png"
stats:
  Years Of Experience: "12"
  Projects Completed: "85"
  Active Projects: "15"
  Happy Customers: "95"
news:
  - img: "/static/img/news/01.png"
    title: "Obtengamos una solución para el trabajo de construcción de edificios"
    date: "22 de junio, 2024"
  - img: "/static/img/news/02.png"
    title: "Obtengamos una solución para el trabajo de construcción de edificios"
    date: "22 de junio, 2024"
  - img: "/static/img/news/03.png"
    title: "Obtengamos una solución para el trabajo de construcción de edificios"
    date: "22 de junio, 2024"
contact:
  cta: "¿Quieres unirte a Interno?"
  desc: "Es un hecho establecido que un lector se distraerá mirando."
  button: "Conéctate con nosotros"
services:
  Kitchen: "Cocina"
  Living Area: "Área de Estar"
  Bathroom: "Baño"
  Bedroom: "Dormitorio"
socials:
  facebook: "https://facebook.com"
  twitter: "https://twitter.com"
  instagram: "https://instagram.com"
year: "2025"
    `

	var result map[string]interface{}
	err := yaml.Unmarshal([]byte(yamlData), &result)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	// Step 2: Convert the map to JSON
	jsonData, err := json.Marshal(result)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	fmt.Println(string(jsonData))
}
