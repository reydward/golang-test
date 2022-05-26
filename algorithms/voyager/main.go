package main

import (
	"fmt"
	"strings"
)

func getTime(s string) int64 {
	input := strings.Split("A"+s, "")
	alphabetSlice := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	alphabet := make(map[string]int)
	minimunTime := 0

	for position, letter := range alphabetSlice {
		alphabet[letter] = position
	}

	for i := 0; i < len(input)-1; i++ {
		minimunTimeInputLetter := alphabet[input[i]] - alphabet[input[i+1]]
		if minimunTimeInputLetter > len(alphabetSlice)/2 {
			minimunTimeInputLetter = len(alphabet) - minimunTimeInputLetter
		}
		minimunTime += minimunTimeInputLetter
	}
	return int64(minimunTime)
}

func main() {
	s := "BZA"
	result := getTime(s)
	fmt.Println(result)
}
