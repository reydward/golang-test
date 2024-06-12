package level1

import (
	"fmt"
	"strconv"
)

func F() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	G(0)
	fmt.Println("Returned normally from g.")
}

func G(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	G(i + 1)
}

func CadenaANumero(cadena string) int {
	defer func() {
		recuperado := recover()
		if recuperado != nil {
			fmt.Println(recuperado)
		}
	}()
	numero, err := strconv.Atoi(cadena)
	if err != nil {
		panic(err)
	}
	return numero
}
