package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hfxbse/dhbw-devops/internal"
	"github.com/hfxbse/dhbw-devops/pkg"
	"log"
	"net/http"
)

func main() {
	var handler = internal.Handler{
		Auth: pkg.Auth{SecretKey: []byte("secret-key")},
		Products: []pkg.Product{
			{ID: 4, Name: "Hammer", Price: 9.99},
			{ID: 5, Name: "Wrench", Price: 14.99},
			{ID: 6, Name: "Screwdriver", Price: 7.99},
			{ID: 7, Name: "Drill", Price: 49.99},
			{ID: 8, Name: "Pliers", Price: 12.99},
		},
	}

	router := mux.NewRouter()
	// Auth Service
	router.HandleFunc("/auth/login", handler.AuthLoginHandler).Methods("POST")
	router.HandleFunc("/auth/logout", internal.AuthLogoutHandler).Methods("POST")
	// Product Service
	router.HandleFunc("/products", handler.ProductListHandler).Methods("GET")
	router.HandleFunc("/products/{id}", handler.ProductDetailHandler).Methods("GET")
	// Checkout Service
	router.HandleFunc("/checkout/placeorder", handler.CheckoutPlaceOrderHandler).Methods("POST")
	port := 8080
	log.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
