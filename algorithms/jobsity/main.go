package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strings"
)

func segment(x int32, space []int32) int32 {
	spaceSize := int32(len(space))

	if spaceSize == 1 {
		return space[0]
	}

	count := int32(0)
	currentMinimun := int32(math.MaxInt32)
	maximunGlobal := int32(-1)

	for i := int32(0); i < spaceSize; i++ {
		if count < x {
			currentMinimun = min(currentMinimun, space[i])
			count++
		} else {
			maximunGlobal = max(maximunGlobal, currentMinimun)
			if space[i-count] == currentMinimun {
				currentMinimun = space[i-count+1]
				j := i - count + 1
				for j <= i {
					currentMinimun = min(currentMinimun, space[j])
					j++
				}
			} else {
				currentMinimun = min(currentMinimun, space[i])
			}
		}
	}
	if maximunGlobal == -1 {
		maximunGlobal = currentMinimun
	}

	return maximunGlobal
}

func main() {
	space := []int32{2, 5, 4, 6, 8}
	x := int32(3)
	result := segment(x, space)

	fmt.Println("result: ", result)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
