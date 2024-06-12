package api

import (
	"encoding/json"
	"fmt"
	"golang-test/live-coding/inventory/structures"
	"io/ioutil"
	"net/http"
	"strconv"
)

func CreateProduct(writer http.ResponseWriter, request *http.Request) {
	var product structures.Product
	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Println("Error")
	}
	json.Unmarshal(requestBody, &product)
	//add to the collection products
	GeneratedID++
	product.ID = strconv.Itoa(GeneratedID)
	Inventory = append(Inventory, product)
	//set http response
	writer.WriteHeader(http.StatusCreated)
	//return json object added
	json.NewEncoder(writer).Encode(product)
}

func GetInventory(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode(Inventory)
}
