package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hfxbse/dhbw-devops/checkout/internal"
	"log"
	"net/http"
	"os"
)

func main() {
	url := os.Getenv("AUTH_SERVICE_URL")
	if len(url) < 1 {
		url = "http://localhost:8080"
	}

	var handler = internal.Checkout{
		AuthService: url,
	}

	router := mux.NewRouter()

	// Checkout Service
	router.HandleFunc("/checkout/placeorder", handler.CheckoutPlaceOrderHandler).Methods("POST")
	port := 8081
	log.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
