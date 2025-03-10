package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hfxbse/dhbw-devops/products/internal"
	"github.com/hfxbse/dhbw-devops/products/pkg"
	"log"
	"net/http"
)

func main() {
	var handler = internal.Products{
		All: []pkg.Product{
			{ID: 4, Name: "Hammer", Price: 9.99},
			{ID: 5, Name: "Wrench", Price: 14.99},
			{ID: 6, Name: "Screwdriver", Price: 7.99},
			{ID: 7, Name: "Drill", Price: 49.99},
			{ID: 8, Name: "Pliers", Price: 12.99},
		},
	}

	router := mux.NewRouter()
	// Product Service
	router.HandleFunc("/products", handler.ProductListHandler).Methods("GET")
	router.HandleFunc("/products/{id}", handler.ProductDetailHandler).Methods("GET")

	port := 8082
	log.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
