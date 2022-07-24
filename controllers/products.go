package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/project/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(wr http.ResponseWriter, r *http.Request) {
	products, err := models.GetProducts()
	if err != nil {
		log.Println("Error getting products: ", err)
	}

	temp.ExecuteTemplate(wr, "Index", products)
}

func New(wr http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(wr, "New", nil)
}

func GetToInsert(wr http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		value := r.FormValue("value")
		amount := r.FormValue("amount")

		valueParsed, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Println("Error parsing value: ", err)
		}

		amountParsed, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Error parsing amount: ", err)
		}

		models.InsertProduct(name, description, valueParsed, amountParsed)
	}
	http.Redirect(wr, r, "/", http.StatusMovedPermanently)
}

func GetToDelete(wr http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	models.DeleteProduct(id)

	http.Redirect(wr, r, "/", http.StatusMovedPermanently)
}

func GetToEdit(wr http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product, err := models.GetProductByID(id)
	if err != nil {
		log.Println("Error getting product: ", err)
	}

	temp.ExecuteTemplate(wr, "Edit", product)
}

func GetToUpdate(wr http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		value := r.FormValue("value")
		amount := r.FormValue("amount")

		valueParsed, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Println("Error parsing value: ", err)
		}

		amountParsed, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Error parsing amount: ", err)
		}

		idParsed, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error parsing id: ", err)
		}

		models.UpdateProduct(idParsed, name, description, valueParsed, amountParsed)
	}
	http.Redirect(wr, r, "/", http.StatusMovedPermanently)
}
