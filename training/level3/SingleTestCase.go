package level3
/*
Statement
Create a Testing function to check the behavior of the following function.
If the function returns a different value from the expected one, return an error specifying the test case.
package main

Topics to Practice:
testing, function, conditions, error
*/

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
