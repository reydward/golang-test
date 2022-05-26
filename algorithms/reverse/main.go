package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}


func (s *Stack) Pop() (string) {
	if s.IsEmpty() {
		return ""
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}
}


func reverseOnlyLetters(S string) string {
	var letters Stack
	chars := []rune(S)
	for _, char := range chars {
		if(unicode.IsLetter(char)){
			letters.Push(string(char))
		}
	}

	var sb strings.Builder
	for _, char := range chars {
		if(unicode.IsLetter(char)){
			sb.WriteString(letters.Pop())
		} else {
			sb.WriteString(string(char))
		}
	}

	return sb.String()
}

func main()  {
	input1 := "a-bC-dEf=ghlj!!"
//	fmt.Scan(&input)

	fmt.Println(reverseOnlyLetters(input1))
}
