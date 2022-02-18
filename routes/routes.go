package routes

import (
	"golang-webapp/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new-product", controllers.NewProduct)
	http.HandleFunc("/insert-product", controllers.InsertProduct)
	http.HandleFunc("/delete-product", controllers.DeleteProduct)
}
