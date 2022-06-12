package main

import (
	"log"
	"net/http"

	"example.com/packages/handlers"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server is running...")

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/api/products",handlers.GetProductsHandler).Methods("GET")
	r.HandleFunc("/api/products/{id}",handlers.GetProductHandler).Methods("GET")
	r.HandleFunc("/api/products",handlers.PostProductHandler).Methods("POST")
	r.HandleFunc("/api/products/{id}",handlers.PutProductHandler).Methods("Put")
	r.HandleFunc("/api/products/{id}",handlers.DeleteProductHandler).Methods("Delete")

	server := &http.Server{
		Addr:		":8080",
		Handler:	r,
	}

	server.ListenAndServe()
	log.Println("Server is stopping...")
}