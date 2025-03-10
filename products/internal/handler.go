package internal

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hfxbse/dhbw-devops/products/pkg"
	"net/http"
	"strconv"
)

type Products struct {
	All []pkg.Product
}

func (p *Products) ProductListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(p.All)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Internal Server Error"}`))
		return
	}
	w.Write(response)
}

func (p *Products) ProductDetailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	productID, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Product ID is missing"}`))
		return
	}

	id, err := strconv.Atoi(productID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Product ID has wrong format"}`))
		return
	}

	product := pkg.FindProductByID(p.All, id)
	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Product not found"}`))
		return
	}

	response, err := json.Marshal(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Internal Server Error"}`))
		return
	}

	w.Write(response)
}
