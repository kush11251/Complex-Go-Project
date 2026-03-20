package controllers

import (
	"example.com/project/pkg/models"
	"example.com/project/pkg/services"
	"encoding/json"
	"net/http"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetUserSvc().GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}