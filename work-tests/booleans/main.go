package main

import 	"fmt"

func main()  {
	expired := []bool{false}

	var AnyVinExpired bool = false
	for _, value := range expired {
		if value {
			AnyVinExpired = true
		}
	}

	fmt.Println("Final value:", AnyVinExpired)
}