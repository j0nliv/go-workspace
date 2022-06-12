package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"example.com/packages/helpers"
	"example.com/packages/models"
	"github.com/gorilla/mux"
)


var productsData = make(map[string]models.Product)
var id int = 0

func PostProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	helpers.CheckError(err)

	product.CreatedOn = time.Now()
	id++
	product.ID = id
	key := strconv.Itoa(id)
	productsData[key] = product

	data,err := json.Marshal(product)
	helpers.CheckError(err)

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	for _,product := range productsData {
		products = append(products, product)
	}

	data, err := json.Marshal(products)
	helpers.CheckError(err)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	for _,prd := range productsData {
		if prd.ID == key {
			product = prd
		}
	}

	data, err := json.Marshal(product)
	helpers.CheckError(err)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func PutProductHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	key := vars["id"]

	var productUp models.Product
	err = json.NewDecoder(r.Body).Decode(&productUp)
	helpers.CheckError(err)

	if _, ok := productsData[key]; ok {
		productUp.ID, _ = strconv.Atoi(key)
		productUp.ChangedOn = time.Now()
		delete(productsData, key)
		productsData[key] = productUp
	} else {
		log.Printf("Value not found!")
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	if _, ok := productsData[key]; ok {
		delete(productsData, key)
	} else {
		log.Printf("Value not found!")
	}
	w.WriteHeader(http.StatusOK)
}