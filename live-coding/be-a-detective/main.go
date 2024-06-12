package main

import (
	"fmt"
	"sort"
)

func getPINs(observed string) []string {
	keypad := map[string][]string{
		"1": {"1", "2", "4"},
		"2": {"1", "2", "3", "5"},
		"3": {"2", "3", "6"},
		"4": {"1", "4", "5", "7"},
		"5": {"2", "4", "5", "6", "8"},
		"6": {"3", "5", "6", "9"},
		"7": {"4", "7", "8"},
		"8": {"0", "5", "7", "8", "9"},
		"9": {"6", "8", "9"},
		"0": {"0", "8"},
	}

	if len(observed) == 1 {
		return keypad[observed]
	}

	prevPINs := getPINs(observed[:len(observed)-1])
	lastDigit := observed[len(observed)-1:]
	var result []string
	for _, prevPIN := range prevPINs {
		if lastDigit == "0" {
			result = append(result, prevPIN+"0")
		} else {
			for _, digit := range keypad[lastDigit] {
				result = append(result, prevPIN+digit)
			}
		}
	}
	sort.Strings(result)

	// Modificación para agregar saltos de línea
	for _, pin := range result {
		fmt.Println(pin)
	}

	return result // Return the original slice (optional)
}

func main() {
	var pin string
	fmt.Scanln(&pin)

	var out_ []string = getPINs(pin)
	for _, str := range out_ {
		fmt.Print(str, "")
	}
}

/*func main() {
	observedPIN := "11"
	possiblePINs := getPINs(observedPIN)
	for _, pin := range possiblePINs {
		fmt.Println(pin)
	}
}*/
