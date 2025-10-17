package main

import "fmt"

func LongestConsecutive(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	numSet := make(map[int]bool)

	for _, num := range arr {
		numSet[num] = true
	}

	maxLength := 0
	var varFiltersCg int

	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentLength := 1

			for numSet[currentNum+1] {
				currentNum++
				currentLength++
			}

			if currentLength > maxLength {
				maxLength = currentLength
				varFiltersCg = maxLength
			}
		}
	}

	_ = varFiltersCg

	return maxLength
}

// readline function for testing
func readline() []int {
	return []int{5, 6, 1, 2, 8, 9, 7}
}

// do not modify below here, readline is our function
// that properly reads in the input for you
func main() {
	fmt.Println(LongestConsecutive(readline()))
}
