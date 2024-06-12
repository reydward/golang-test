/*
Problem:
- Merge two sorted arrays.

Example:
- Input: []int{1, 3, 5}, []int{2, 4, 6}
  Output: []int{1, 2, 3, 4, 5, 6}
- Input: []int{1, 3, 5}, []int{2, 4, 6, 7}
  Output: []int{1, 2, 3, 4, 5, 6, 7},

Approach:
- Since these arrays are sorted, can use two pointers approach to iterate
  through both of them and append the smaller value to a new merged list at
  each step.

Solution:
- Have two pointers start at the beginning of these two arrays.
- While both of them does not reach the end, compare two current values
  at each step and append the smaller one two a new merged list.
- Move the two pointers up accordingly as values get merged in.
- In the case where one of these pointers reach the end first and the
  other one is still in the middle of the array, simply add the rest of
  its values to the merged list since they are all sorted and guaranteed
  to be in ascending order.

Cost:
- O(n) time, O(n) space.
*/

package main

import "fmt"

func mergeSortedArrays(arr1, arr2 []int) []int {
	// Inicializar las variables para los índices de arr1 y arr2
	i, j := 0, 0
	// Inicializar una lista para el resultado
	merged := make([]int, 0)

	// Mientras ambos índices estén dentro del rango de las listas
	for i < len(arr1) && j < len(arr2) {
		// Comparar los valores actuales de las dos listas
		if arr1[i] <= arr2[j] {
			// Si el valor de arr1 es menor o igual, agregarlo al resultado
			merged = append(merged, arr1[i])
			// Mover el índice de arr1 hacia adelante
			i++
		} else {
			// Si el valor de arr2 es menor, agregarlo al resultado
			merged = append(merged, arr2[j])
			// Mover el índice de arr2 hacia adelante
			j++
		}
	}

	// Agregar los elementos restantes de arr1 (si los hay)
	for i < len(arr1) {
		merged = append(merged, arr1[i])
		i++
	}

	// Agregar los elementos restantes de arr2 (si los hay)
	for j < len(arr2) {
		merged = append(merged, arr2[j])
		j++
	}

	return merged
}

func main() {
	// Ejemplos de arrays ordenados
	arr1 := []int{1, 3, 5}
	arr2 := []int{2, 4, 6}

	// Combinar los dos arrays
	result := mergeSortedArrays(arr1, arr2)

	// Imprimir el resultado
	fmt.Println(result) // Output: [1 2 3 4 5 6]
}
