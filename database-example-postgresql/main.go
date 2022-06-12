package main

import (
	"example.com/postgrego/models"
)

func main() {

	product := models.Product{
		Title:			"Toy Car",
		Description: 	"1:32 Toy Car",
		Price:			255.99,
	}
	models.InsertProduct(product)

	/*
	upProd := models.Product{
		ID:				2
		Title:			"Toy Car",
		Description: 	"1:32 Toy Car",
		Price:			255.99,
	}
	models.UpdateProduct(product)
	

	models.GetProducts()

	models.GetProdByID(2)

	*/
	
}