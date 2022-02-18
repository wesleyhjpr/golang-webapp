package models

import (
	"golang-webapp/db"

	_ "github.com/lib/pq"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() []Product {
	db := db.ConnectDataBase()

	allProducts, err := db.Query("SELECT * FROM products")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for allProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = allProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	defer db.Close()

	return products
}

func CreateProduct(name, description string, quantity int, price float64) {
	db := db.ConnectDataBase()

	insertProductDB, err := db.Prepare("INSERT INTO products (name, description, quantity, price) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertProductDB.Exec(name, description, quantity, price)

	defer db.Close()
}

func DeleteProduct(idProduct string) {
	db := db.ConnectDataBase()

	deleteProduct, err := db.Prepare("DELETE FROM products WHERE id=$1")

	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(idProduct)

	defer db.Close()
}
