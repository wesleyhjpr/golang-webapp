package main

import (
	"net/http"
	"text/template"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Amount      int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{"T-shirt", "Very pretty", 39, 5},
		{"Tenis", "Confortavel", 89, 3},
		{"Head Phone", "very good", 59, 2},
		{"New Product", "Nice", 1.99, 1},
	}

	templates.ExecuteTemplate(w, "Index", products)
}
