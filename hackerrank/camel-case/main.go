/*
There is a sequence of words in CamelCase as a string of letters, , having the following properties:

It is a concatenation of one or more words consisting of English letters.
All letters in the first word are lowercase.
For each of the subsequent words, the first letter is uppercase and rest of the letters are lowercase.
Given , determine the number of words in .

Example
s = oneTwoThree
There are  words in the string: 'one', 'Two', 'Three'.
Solution: We initialize a counter count with 1 since the first word always starts with a lowercase letter.
We iterate over each character in the string s.
For each character, if it is an uppercase letter (checked using unicode.IsUpper), we increment the counter count.
Finally, we return the counter count which represents the number of words in the CamelCase string.
*/
package main

import (
	"unicode"
)

func camelcase(s string) int32 {
	var count int32 = 1 // Start with 1 because the first word is lowercase
	for _, char := range s {
		if unicode.IsUpper(char) {
			count++
		}
	}
	return count
}

func main() {
	//s := "oneTwoThree"
	s := "saveChangesInTheEditor"
	result := camelcase(s)
	println(result)
}

/*
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := camelcase(s)

	fmt.Fprintln(writer, result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err != nil {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
*/
