package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// Product will storage the information related to a Product
type Product struct {
	ID    string  `json:"id,omitempty"`
	Code  string  `json:"code,omitempty"`
	Name  string  `json:"name,omitempty"`
	Price float64 `json:"price,omitempty"`
}

type allProducts []Product

var products = allProducts{
	{
		ID: "1",
		Code: "Code 1",
		Name: "Product 1",
		Price: 1,
	},
}

var generatedID = 1

func main()  {
	router := mux.NewRouter()
	router.HandleFunc("/products", getAll).Methods("GET")
	router.HandleFunc("/products/{id}", getProduct).Methods("GET")
	router.HandleFunc("/products", createProduct).Methods("POST")
	router.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func deleteProduct(writer http.ResponseWriter, request *http.Request) {
	productID := mux.Vars(request)["id"]
	for i, product := range products {
		if product.ID == productID {
			products = append(products[:i], products[i+1:]...)
		}
	}
}

func updateProduct(writer http.ResponseWriter, request *http.Request) {
	productID := mux.Vars(request)["id"]
	var newProduct Product
	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil{
		fmt.Println("Error")
	}
	json.Unmarshal(requestBody, &newProduct)

	for i, product := range products {
		if product.ID == productID {
			newProduct.ID = productID
			products[i] = newProduct
			products = append(products[:i], product)
			json.NewEncoder(writer).Encode(newProduct)
		}
	}

}

func createProduct(writer http.ResponseWriter, request *http.Request) {
	//receive the object
	var product Product
	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil{
		fmt.Println("Error")
	}
	json.Unmarshal(requestBody, &product)
	//add to the collection products
	generatedID++
	product.ID = strconv.Itoa(generatedID)
	products = append(products, product)
	//set http response
	writer.WriteHeader(http.StatusCreated)
	//return json object added
	json.NewEncoder(writer).Encode(product)
}

func getProduct(writer http.ResponseWriter, request *http.Request) {
	productID := mux.Vars(request)["id"]

	for _, product := range products {
		if product.ID == productID {
			json.NewEncoder(writer).Encode(product)
		}
	}
}

func getAll(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode(products)
}
