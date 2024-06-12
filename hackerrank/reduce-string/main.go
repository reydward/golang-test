/*
Reduce a string of lowercase characters in range ascii[‘a’..’z’]by doing a series of operations. In each operation,
select a pair of adjacent letters that match, and delete them.
Delete as many characters as possible using this method and return the resulting string. If the final string is empty, return Empty String

Example.
s = 'aab'
aab shortens to b in one operation: remove the adjacent a characters.

s = 'abba'
Remove the two 'b' characters leaving 'aa'. Remove the two 'a' characters to leave ”. Return 'Empty String'.

Solution: We initialize an empty stack (slice of runes).
We iterate over each character in the input string.
For each character:
If the stack is not empty and the top of the stack is the same as the current character, we pop the stack (remove the top element).
Otherwise, we push the current character onto the stack.
After processing all characters, if the stack is empty, we return "Empty String".
Otherwise, we convert the stack back to a string and return it.
*/
package main

import "fmt"

func superReducedString(s string) string {
	var stack []rune
	for _, char := range s {
		if len(stack) > 0 && stack[len(stack)-1] == char {
			stack = stack[:len(stack)-1] // Pop from stack
		} else {
			stack = append(stack, char) // Push to stack
		}
		fmt.Println("stack:", stack)
	}
	if len(stack) == 0 {
		return "Empty String"
	}
	return string(stack)
}

func main() {
	//	s := "aab"
	s := "aaabccddd"
	fmt.Println("s:", s)
	result := superReducedString(s)
	fmt.Println(result)
}

/*
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := superReducedString(s)

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
