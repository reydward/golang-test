package main

import "fmt"

func modifyValueByValue(x int) {
	x = x * 2 // Modificar la copia de x
	fmt.Println("Dentro de la función:", x)
}
func modifyValueByReference(x *int) {
	*x = *x * 2 // Modificar el valor apuntado por x
	fmt.Println("Dentro de la función:", *x)
}
func main() {
	// Paso por valor
	value := 10
	modifyValueByValue(value)
	fmt.Println("Fuera de la función:", value) // value sigue siendo 10

	// Paso por referencia
	modifyValueByReference(&value)
	fmt.Println("Fuera de la función:", value) // value se modifica a 20
}
