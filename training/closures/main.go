package main

import "fmt"

func main() {
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())

	// Definimos una función que retorna una clausura
	multiply := multiplier(2)
	fmt.Println("Multiplicación:", multiply(4)) // Salida: Multiplicación: 8

}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

// Función que retorna una clausura
func multiplier(factor int) func(int) int {
	// Retorna una clausura que multiplica un número por el factor dado
	return func(x int) int {
		println("X:", x)
		println("Factor:", factor)
		return x * factor
	}
}
