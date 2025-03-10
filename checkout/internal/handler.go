package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Checkout struct {
	AuthService string
}
type VerificationResult struct {
	Valid bool
}

func (c *Checkout) CheckoutPlaceOrderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := r.Header.Get("Authorization")

	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "Missing authorization header"}`))
		return
	}

	response, err := http.Post(
		fmt.Sprintf(`%s/auth/verify`, c.AuthService),
		"application/json",
		bytes.NewBuffer([]byte(fmt.Sprintf(`{"token": "%s"}`, token))),
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Authentication failed"}`))
		return
	}

	var result VerificationResult
	err = json.NewDecoder(response.Body).Decode(&result)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Authentication failed"}`))
		return
	}

	if result.Valid {
		// In this simple example, we'll just return a success message
		w.Write([]byte(`{"message": "Order placed successfully"}`))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "Invalid token"}`))
	}
}
