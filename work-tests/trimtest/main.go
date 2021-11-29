package main

import (
	"fmt"
	"strings"
)

func main()  {
	somestring := " my string "
	trimstring := strings.Trim(somestring, " ")

	fmt.Println("start"+trimstring+"end")
}
