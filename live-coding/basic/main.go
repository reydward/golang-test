package main

import "fmt"

type MyStruct struct {
	name    string
	surname string
}

func main() {
	stringArray := []string{"one", "two"}
	myStruct := MyStruct{"Eduard", "Reyes"}
	var myStructArray []MyStruct

	for _, value := range stringArray {
		fmt.Println(value)
	}

	myStructArray = append(myStructArray, myStruct)
	for _, value := range myStructArray {
		fmt.Println(value)
	}

	fmt.Println(myStruct)
}
