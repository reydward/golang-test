package level3

import (
	"strconv"
	"testing"
)

/*
Statement
Create a Testing function to check the behavior of the following function.
If the function returns a different value from the expected one, return an error specifying the test case.
package main

Topics to Practice:
testing, function, conditions, error
*/

func TestIntMin(t *testing.T) {
	var result int
	type test struct {
		data []int
		result int
	}

	tests := []test{
		{[]int{1, 2}, 1},
		{[]int{2, 3}, 2},
		{[]int{3, 3}, 3},
		{[]int{4, 5}, 4},
	}

	for _, v := range tests{
		result = IntMin(v.data[0], v.data[1])
		if result != v.result {
			t.Error("Expected " + strconv.Itoa(v.result) + " but got ", result)
		}
	}
}
