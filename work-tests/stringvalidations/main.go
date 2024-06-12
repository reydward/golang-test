package main

import "fmt"

func main() {
	var mystr *string //initialize a string
	mystr = ""

	if mystr == nil {
		fmt.Println("String is nil") //if the length of string is 0 print its empty
	}

	if len(*mystr) == 0 {
		fmt.Println("String is empty") //if the length of string is 0 print its empty
	} else {
		fmt.Println("String is not empty") //else print not empty
	}
}
