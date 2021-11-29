package main

import (
	"fmt"
	"golang-test/level1"
)

func main() {

	//Age Filter
	agesSlice := []int{1, 2, 3, 4, 5}
	fmt.Print(level1.AgeFilter(agesSlice, 1, 40))

/*
	//Simple Calculator Operations
	fmt.Print(level1.Add(2, 3).Result)
	fmt.Print(level1.Substract(2, 3).Result)
	fmt.Print(level1.Multiply(2, 3).Result)
	fmt.Print(level1.Divide(2, 0).Result)
 */

//	fmt.Println(level1.C())

/*
	//Testing handling errors
	level1.F()
	fmt.Println("Returned normally from f.")*/

/*
	//Testing handling errors
	fmt.Println("Inicia")           // Inicia
	fmt.Println(level1.CadenaANumero("2")) // 2
	// strconv.Atoi: parsing "j": invalid syntax
	fmt.Println(level1.CadenaANumero("j")) // 0
	fmt.Println("Fin")              // Fin

	product := level1.Product{
		ID: 1,
		Name: "Name 1",
	}

	level1.Ad(product)*/
}

