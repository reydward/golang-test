package main

import (
	"fmt"
	"strings"
)

func main()  {
	brandApi := []string{"one,two"}

	brand := strings.Join(brandApi, ",")

	fmt.Println(brand)
}
