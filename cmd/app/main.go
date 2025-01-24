package main

import (
	"fmt"
	"net/http"

	"cerescms.net/internal/controllers"
)

func main() {
	fmt.Println("Server is running on http://localhost:4000")
	http.ListenAndServe(":4000", controllers.Router())
}
