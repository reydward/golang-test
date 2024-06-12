/*
The website considers a password to be strong if it satisfies the following criteria:

Its length is at least .
It contains at least one digit.
It contains at least one lowercase English character.
It contains at least one uppercase English character.
It contains at least one special character. The special characters are: !@#$%^&*()-+
She typed a random string of length  in the password field but wasn't sure if it was strong. Given the string she typed, can you find the minimum number of characters she must add to make her password strong?

Note: Here's the set of types of characters in a form you can paste in your solution:

numbers = "0123456789"
lower_case = "abcdefghijklmnopqrstuvwxyz"
upper_case = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
special_characters = "!@#$%^&*()-+"
Example
pasword = "2bbbb"
This password is 5 characters long and is missing an uppercase and a special character. The minimum number of characters to add is .

pasword = "2bb#A"
This password is 5 characters long and has at least one of each character type. The minimum number of characters to add is .

Solution:
*/
package main

import (
	"fmt"
	"strings"
)

func minimumNumber(n int32, password string) int32 {
	numbers := "0123456789"
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharacters := "!@#$%^&*()-+"

	missingCount := 0
	hasNumber := false
	hasLower := false
	hasUpper := false
	hasSpecial := false

	for _, char := range password {
		switch {
		case strings.ContainsRune(numbers, char):
			hasNumber = true
		case strings.ContainsRune(lowerCase, char):
			hasLower = true
		case strings.ContainsRune(upperCase, char):
			hasUpper = true
		case strings.ContainsRune(specialCharacters, char):
			hasSpecial = true
		}
	}

	if !hasNumber {
		missingCount++
	}
	if !hasLower {
		missingCount++
	}
	if !hasUpper {
		missingCount++
	}
	if !hasSpecial {
		missingCount++
	}

	if len(password)+missingCount < 6 {
		return 6 - int32(len(password))
	}

	return int32(missingCount)
}

func main() {
	n := int32(5)
	password := "2abcde"
	result := minimumNumber(n, password)
	fmt.Println(result)
}
