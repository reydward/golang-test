package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func main() {

	// Variable declaration
	var count int = 10
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(count, i, j, k, c, python, java)

	// Creación de un slice de enteros
	numeros := []int{5, 12, 8, 23, 42}

	// Iteración con for range
	for indice, valor := range numeros {
		fmt.Printf("Índice: %d, Valor: %d\n", indice, valor)
	}

	// Iteración solo sobre los valores (ignorando los índices)
	for _, valor := range numeros {
		fmt.Println("Valor:", valor)
	}

	// Array
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// For loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for key, value := range primes {
		fmt.Println(key, value)
	}

	// Slice
	var s []int = primes[1:4]
	fmt.Println(s)

	// Map
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// If statement with error control
	/*	if err != nil {
			log.Println(err, "Failed to retrieve inventory")
			return nil, err
		}

		if err := request.ParseForm(); err != nil {
			log.Error(err, "Error parsing form")
			return err
		}*/

	// Struct
	type Package struct {
		ID              int     `json:"id"`
		OrderID         int     `json:"orderId"`
		VendorID        int     `json:"vendorId"`
		PackageVendorID *string `json:"packageVendorId"`
	}

	// Function
	fmt.Println(add(42, 13))
	// Method
	// TO IMPLEMENT WITH receiver

}
