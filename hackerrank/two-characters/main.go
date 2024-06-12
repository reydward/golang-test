/*
Given a string, remove characters until the string is made up of any two alternating characters. When you choose a character to remove,
all instances of that character must be removed. Determine the longest string possible that contains just two alternating letters.

Example
s = 'abaacdabd'
Delete a, to leave bcdbd. Now, remove the character c to leave the valid string bdbd with a length of 4.
Removing either b or d at any point would not result in a valid string. Return 4.
Given a string s, convert it to the longest possible string t made up only of alternating characters.
Return the length of string t. If no string t can be formed, return 4.

Solution:
Character Counts:
We use a map (charCount) to store the unique characters in the string along with their frequencies.
Nested Loops for Pairs: We use two nested loops to iterate over all possible pairs of unique characters (char1 and char2) from the map.
We skip pairs where the characters are the same or where count1 is greater than count2 to avoid redundant checks.
Alternation Check:
We initialize valid to true and lastChar to a dummy value.
We iterate through the original string (s).
For each character, we check if it's one of the two selected characters.
If it is, and it's different from the previous character (lastChar), it's a valid part of an alternating string. Otherwise, the string is invalid.
We update lastChar for the next iteration.
Update Max Length:
If the string formed by char1 and char2 is valid and alternating, we calculate its length (count1 + count2) and update maxLen if it's the longest valid string found so far.
Return Max Length:
After checking all pairs, we return the maximum length of a valid alternating string found, or 0 if no such string exists.
Helper Function max: A simple helper function is included to find the maximum of two integers.
*/
package main

import (
	"fmt"
)

func alternate(s string) int32 {
	// Get unique characters and their counts
	charCount := make(map[rune]int)
	for _, char := range s {
		charCount[char]++
	}

	var maxLen int32 = 0
	for char1, count1 := range charCount {
		for char2, count2 := range charCount {
			// Skip the same character and ensure count1 <= count2
			if char1 == char2 || count1 > count2 {
				continue
			}

			// Check if the string with these two characters is alternating
			valid := true
			lastChar := rune(0)
			for _, char := range s {
				if char != char1 && char != char2 {
					continue
				}
				if char == lastChar {
					valid = false
					break
				}
				lastChar = char
			}

			if valid {
				maxLen = max(maxLen, int32(count1+count2))
			}
		}
	}
	return maxLen
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "beabeefeab"
	result := alternate(s)
	fmt.Println(result)
}
