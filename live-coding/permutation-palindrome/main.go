/*
Problem:
- Given a string, check if its permutation is a palindrome.

Example:
- Input: "ivicc"
  Output: true
- Input: "civic"
  Output: true

Approach:
- To determine if a permutation is a palindrome, need to check if each
  character in the string appears an even number of times, allowing for
  only one character to appear an odd time, that is the middle one.
- Could use a hashmap store the characters and their number of occurrences.

Solution:
- As we iterate through the string, use a hashmap to add a character if
  we haven't seen it and remove it if it's already there.
- After the iteration, if we're left with less or equal than a character in
  the map, we have a palindrome.

Cost:
- O(n) time, O(1) space.
- The space complexity is O(n) due to the hashmap, but since there are
  only a constant number of characters in Unicode, we could treat it
  as O(1).
*/
package main

import "fmt"

func isPalindromePermutation(s string) bool {
	charCounts := map[rune]int{} // Mapa para contar caracteres (runas)

	for _, char := range s {
		charCounts[char]++         // Incrementar el contador del caracter
		if charCounts[char] == 2 { // Si aparece por segunda vez...
			delete(charCounts, char) // Eliminarlo de la cuenta
		}
	}

	return len(charCounts) <= 1 // Condición de palíndromo
}

func main() {
	fmt.Println(isPalindromePermutation("ivicc")) // Output: true
	fmt.Println(isPalindromePermutation("civic")) // Output: true
	fmt.Println(isPalindromePermutation("hello")) // Output: false
}
