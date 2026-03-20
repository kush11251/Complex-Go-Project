package cmd

import (
	"example.com/project/pkg/controllers"
	"net/http"
)

func Router() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("/users", controllers.GetUserHandler)
	r.HandleFunc("/products", controllers.GetProductHandler)
	return r
}