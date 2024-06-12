/*
Insert element into sorted list
Given a sorted list with an unsorted number  in the rightmost cell, can you write some simple code to insert  into the array so that it remains sorted?
Since this is a learning exercise, it won't be the most efficient way of performing the insertion. It will instead demonstrate the brute-force method in detail.
Assume you are given the array  indexed . Store the value of . Now test lower index values successively from  to  until you reach a value that is lower than, at  in this case.
Each time your test fails, copy the value at the lower index to the current index and print your array.
When the next lower indexed value is smaller than , insert the stored value at the current index and print the entire array.

Example
5
2 4 6 8 3
Sample Output
2 4 6 8 8
2 4 6 6 8
2 4 4 6 8
2 3 4 6 8

Solution: Store the Target: The last element of the array arr[n-1] is the unsorted element to be inserted. We store it in the target variable.

Iterate Backwards: We start a loop from the second-to-last element (i = n-2) and iterate backwards as long as:
We haven't reached the beginning of the array (i >= 0).
The current element arr[i] is greater than the target.
Shift Larger Elements: Inside the loop, if the current element is larger than the target, we shift it one position to the right (arr[i+1] = arr[i]). This creates space for the target to be inserted.
Print Intermediate Steps: After each shift, we print the current state of the array using fmt.Println(arr). This fulfills the challenge's requirement to print the array after every shift or insertion.
Insert the Target:  Once we find the correct position (either because we reached the beginning or found a smaller element), we insert the target element at arr[i+1].
Print Final Array: Finally, we print the sorted array after the insertion is complete using fmt.Println(arr).
*/
package main

import (
	"fmt"
)

func insertionSort1(n int32, arr []int32) {
	target := arr[n-1] // Store the unsorted element
	i := n - 2         // Start at the index before the unsorted element

	for ; i >= 0 && arr[i] > target; i-- {
		arr[i+1] = arr[i] // Shift larger elements to the right
		printArray(arr)   // Print the array after each shift
	}

	arr[i+1] = target // Insert the target into its correct position
	printArray(arr)
	//fmt.Println(arr)  // Print the final sorted array
}

func printArray(arr []int32) {
	for i := range arr {
		if i == len(arr)-1 {
			fmt.Print(arr[i], "\n")
		} else {
			fmt.Print(arr[i], " ")
		}
	}
}

func main() {
	n := int32(5)
	arr := []int32{2, 4, 6, 8, 3}
	insertionSort1(n, arr)
}
