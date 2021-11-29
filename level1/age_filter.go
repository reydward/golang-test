/*
Age Filter
Statement
Create a function that will filter a Slice of ages that are between the range.
The function will receive two numbers and a slice of ages as parameters. It should return the ages between the range

Topics to Practice:
slice, append, function, for loop, control flow, return value
*/
package level1

func AgeFilter(slice []int, age1 int, age2 int) (filteredSlice []int) {
	filteredSlicex := []int{}
	for v := range slice {
		if v >= age1 && v <= age2 {
			filteredSlicex = append(filteredSlicex, v)
		}
	}

	return filteredSlicex
}
