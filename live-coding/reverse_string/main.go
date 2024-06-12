package main

import (
	"fmt"
)

/*
Problem:
- Given a list of string, reverse its order in place.

Example:
- Input: []string{"a", "b", "c", "d"}
  Output: []string{"d", "c", "b", "a"}

Approach:
- Use two pointers approach to swap two values on both ends as we move toward
  the middle.

Solution:
- Initialize the two pointers, one starts at the beginning and one starts at
  the end of the list.
- While the start pointer does not meet the end pointer in the middle, keep
  swapping these two values.
- Move the start pointer up and move the end pointer down.

Cost:
- O(n) time, O(1) space.
*/

func reverseStringList(lst []string) {
	// Inicializar los punteros de inicio y final
	start := 0
	end := len(lst) - 1

	// Mientras los punteros no se crucen
	for start < end {
		// Intercambiar los elementos en los extremos
		lst[start], lst[end] = lst[end], lst[start]
		// Mover el puntero de inicio hacia adelante
		start++
		// Mover el puntero de final hacia atrás
		end--
	}
}

func main() {
	// Lista de strings de ejemplo
	lst := []string{"a", "b", "c", "d"}

	// Imprimir la lista original
	fmt.Println("Original:", lst)

	// Invertir la lista en su lugar
	reverseStringList(lst)

	// Imprimir la lista invertida
	fmt.Println("Reversed:", lst)
}
