package routes

import (
	"net/http"

	"github.com/project/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.GetToInsert)
	http.HandleFunc("/delete", controllers.GetToDelete)
	http.HandleFunc("/edit", controllers.GetToEdit)
	http.HandleFunc("/update", controllers.GetToUpdate)
}
