package main

import "fmt"

func main() {
	// Crear un slice de enteros
	var numbers []int
	numbers = []int{1, 2, 3, 4, 5}

	// Agregar elementos al slice usando la función append
	numbers = append(numbers, 6)
	numbers = append(numbers, 7)

	fmt.Println("Slice después de agregar elementos:", numbers)

	// Dividir el slice
	middleIndex := len(numbers) / 2
	firstHalf := numbers[:middleIndex]
	secondHalf := numbers[middleIndex:]

	fmt.Println("Primera mitad del slice:", firstHalf)
	fmt.Println("Segunda mitad del slice:", secondHalf)

	// Eliminar un elemento del slice
	indexToDelete := 2
	numbers = append(numbers[:indexToDelete], numbers[indexToDelete+1:]...)

	fmt.Println("Slice después de eliminar el elemento en la posición", indexToDelete, ":", numbers)

	// Crear un slice multidimensional de strings
	var matrix [][]string

	// Agregar elementos al slice
	matrix = append(matrix, []string{"a", "b", "c"})
	matrix = append(matrix, []string{"d", "e", "f"})
	matrix = append(matrix, []string{"g", "h", "i"})

	// Imprimir el slice
	fmt.Println("Slice multidimensional de strings:", matrix)
}
