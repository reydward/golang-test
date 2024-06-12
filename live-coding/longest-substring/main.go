/*
Problem:
- Given a string, find the length of the longest substring without repeating characters.

Example:
- Input: "abcabcbb"
  Output: 3
  Explanation: The longest substring is "abc" with the length of 3.
- Input: "bbbbb"
  Output: 1
  Explanation: The longest substring is "b" with the length of 1.

Approach:
- Iterate through the string and keep track of the maximum length of non-repeating
  characters using a hashmap that maps characters to their indices.
- Could skip characters immediately if we found a repeating character.

Solution:
- Initialize a map that maps characters to their indices.
- Initialize a start index and end index to keep track of the start and end of
  a substring.
- Iterate through the string and check if we have seen the current character
  before in the map.
- If so, update the start index.
- Otherwise, cache the current index and update the maximum length if we found
  a larger one.
- Return the maximum length in the end.

Cost:
- O(n) time, O(m) cost where m < n and  n is the length of the string.
*/

package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	// Crear un mapa para realizar un seguimiento de los índices de los caracteres
	charMap := make(map[byte]int)
	maxLength := 0
	start := 0

	// Iterar a través de la cadena utilizando un enfoque de ventana deslizante
	for end := 0; end < len(s); end++ {
		// Si el carácter está en el mapa y su índice es mayor o igual que el inicio actual,
		// actualizamos el inicio al índice siguiente del carácter repetido
		_, ok := charMap[s[end]]

		//if _, ok := charMap[s[end]]; ok && charMap[s[end]] >= start {
		if ok && charMap[s[end]] >= start {
			start = charMap[s[end]] + 1
		}
		// Actualizamos el índice del carácter en el mapa
		charMap[s[end]] = end
		// Actualizamos la longitud máxima de la subcadena sin repetir
		maxLength = max(maxLength, end-start+1)
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // Output: 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // Output: 1
}
