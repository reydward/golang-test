package main

import "fmt"

func main()  {
	isOrderIntegrated := false
	isOrderCanceled := false
	isOrderError := true

	if isOrderError {
		fmt.Println("isOrderError")
	} else if isOrderIntegrated {
		fmt.Println("isOrderIntegrated")
	} else if isOrderCanceled {
		fmt.Println("isOrderCanceled")
	}

}

