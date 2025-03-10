package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hfxbse/dhbw-devops/auth/internal"
	"github.com/hfxbse/dhbw-devops/auth/pkg"
	"log"
	"net/http"
)

func main() {
	var handler = internal.Auth{
		Provider: pkg.Auth{SecretKey: []byte("secret-key")},
	}

	router := mux.NewRouter()
	// Auth Service
	router.HandleFunc("/auth/login", handler.AuthLoginHandler).Methods("POST")
	router.HandleFunc("/auth/logout", internal.AuthLogoutHandler).Methods("POST")
	router.HandleFunc("/auth/verify", handler.AuthVerifyHandler).Methods("POST")

	port := 8080
	log.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
