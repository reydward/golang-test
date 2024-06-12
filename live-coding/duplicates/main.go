/*
Problem:
  - Given an array containing n numbers taken from the range 1 to n. It can
    have some duplicates. Find all those numbers.

Example:
  - Input: []int{5, 4, 7, 2, 3, 5, 3}
    Output: []int{3, 5}

Approach:
  - Similar to missing number problem, can rearrange the array using cyclic
    sort.
  - Those that do not have the correct indices are the duplicate ones.

Cost:
- O(n) time, O(1) space.
*/
package main

import "fmt"

func findDuplicates(nums []int) []int {
	i := 0
	duplicates := []int{}

	for i < len(nums) {
		correctIndex := nums[i] - 1 // Índice correcto para el valor actual
		fmt.Println("i:", i, "correctIndex:", correctIndex, "nums:", nums)
		if nums[i] != nums[correctIndex] {
			nums[i], nums[correctIndex] = nums[correctIndex], nums[i] // Intercambio
		} else {
			i++ // Continuar sólo si el elemento está en su lugar correcto
		}
	}

	// Los números que no están en su lugar correcto son duplicados
	for i, num := range nums {
		if num != i+1 {
			duplicates = append(duplicates, num)
		}
	}

	return duplicates
}

func main() {
	nums := []int{5, 4, 7, 2, 3, 5, 3}
	duplicates := findDuplicates(nums)
	fmt.Println("Números duplicados:", duplicates)
}
