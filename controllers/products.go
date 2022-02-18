package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"golang-webapp/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAllProducts()

	templates.ExecuteTemplate(w, "Index", allProducts)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New-product", nil)
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		quantity := r.FormValue("quantity")
		price := r.FormValue("price")

		priceConverted, err := strconv.ParseFloat(price, 64)

		if err != nil {
			log.Println("Error to convert price:", err)
		}

		quantityConverted, err := strconv.Atoi(quantity)

		if err != nil {
			log.Println("Error to convert quantity:", err)
		}

		models.CreateProduct(name, description, quantityConverted, priceConverted)
	}
	http.Redirect(w, r, "/", 301)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")

	models.DeleteProduct(idProduct)

	http.Redirect(w, r, "/", 301)
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")

	product := models.GetProductById(idProduct)

	templates.ExecuteTemplate(w, "Edit-product", product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		quantity := r.FormValue("quantity")
		price := r.FormValue("price")

		idConverted, err := strconv.Atoi(id)

		if err != nil {
			log.Println("Error to convert id:", err)
		}

		priceConverted, err := strconv.ParseFloat(price, 64)

		if err != nil {
			log.Println("Error to convert price:", err)
		}

		quantityConverted, err := strconv.Atoi(quantity)

		if err != nil {
			log.Println("Error to convert quantity:", err)
		}

		models.UpdateProduct(idConverted, name, description, quantityConverted, priceConverted)
	}
	http.Redirect(w, r, "/", 301)
}
