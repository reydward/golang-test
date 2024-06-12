/*
Consider an array of numeric strings where each string is a positive number with anywhere from  to  digits. Sort the array's elements in non-decreasing, or ascending order of their integer values and return the sorted array.
Example
unsorted = ['1', '200', '150', '3'].
Return the array ['1', '3', '150', '200'].

Solution: We use sort.Slice to sort the unsorted slice of strings.
In the custom comparator function:
We first compare the lengths of the strings.
If the lengths are different, the shorter string is considered smaller.
If the lengths are the same, we compare the strings lexicographically.
*/
package main

import (
	"fmt"
	"sort"
)

func bigSorting(unsorted []string) []string {
	// Define a custom sorter
	sort.Slice(unsorted, func(i, j int) bool {
		// First compare by length
		if len(unsorted[i]) != len(unsorted[j]) {
			return len(unsorted[i]) < len(unsorted[j])
		}
		// If lengths are the same, compare lexicographically
		return unsorted[i] < unsorted[j]
	})
	return unsorted
}

func main() {
	unsortedArray := []string{"6", "31415926535897932384626433832795", "1", "3", "10", "3", "5"}
	sortedArray := bigSorting(unsortedArray)
	fmt.Println(sortedArray)
}

/*
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	var n int
	fmt.Fscanln(reader, &n)

	var unsorted []string

	for i := 0; i < n; i++ {
		var unsortedItem string
		unsortedItem, err = reader.ReadString('\n')
		checkError(err)
		unsorted = append(unsorted, strings.TrimSpace(unsortedItem))
	}

	result := bigSorting(unsorted)

	for _, resultItem := range result {
		fmt.Fprintln(writer, resultItem)
	}

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
*/
