package main

import (
	"fmt"
	"net/http"

	"cerescms.hapaxredux.net/internal/controllers"
)

func main() {
	fmt.Println("Server is running on http://localhost:4000")
	router := controllers.Router()
	http.ListenAndServe(":4000", router)
}
