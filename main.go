package main

import (
	"net/http"

	"github.com/project/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
