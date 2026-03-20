package controllers

import (
	"example.com/project/pkg/models"
	"example.com/project/pkg/services"
	"encoding/json"
	"net/http"
)

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	products, err := services.GetProductSvc().GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}