/**
Desarrolle un sistema de gestión de inventario concurrente que utiliza una API REST para administrar productos y sus cantidades en stock.
*/

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"golang-test/live-coding/inventory/api"
	"golang-test/live-coding/inventory/structures"
	"net/http"
)

var Inventory []structures.Product
var GeneratedID = 1

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("The server is working"))
}

func main() {
	fmt.Println("Hello, World!")

	router := mux.NewRouter()
	router.HandleFunc("/health", healthCheck).Methods("GET")
	router.HandleFunc("/product", api.CreateProduct).Methods("GET")
	router.HandleFunc("/inventory", api.GetInventory).Methods("GET")

	fmt.Println("Server started on port 8000")
	http.ListenAndServe(":8000", router)
}
