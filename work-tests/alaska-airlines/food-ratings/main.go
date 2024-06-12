package main

import "fmt"

/*
Food ratings
A restaurant has launched an app for their food delivery. ‘n’ reviews about various of their menu items have been submitted.
Users have submitted a rating out of 5 (1 being worst and 5 being best). The manager wants to know the most loved dish in the restaurant.
Find out the dish with the highest average rating.

Note: If two dishes are rated the same, return the dish with the smallest ID.

Sample input
4
512 3
512 5
987 4
123 5

Sample output
123
*/

func mostLovedDish(reviews [][]int) [2]int {
	mostLovedDish := [2]int{}
	dishesMap := make(map[int]int)
	var maxKey, maxValue int

	// Store every dish in a map, in order to have one pair per dish
	// Iterate the map, having an average rate for every dish
	for _, review := range reviews {
		fmt.Println("Review:", review[0])
		if dishesMap[review[0]] == 0 {
			dishesMap[review[0]] += review[1]
		} else {
			dishesMap[review[0]] += review[1] / 2
		}
	}

	// I have the map wih every dish and its corresponding rate average
	// Get the dish with the largest average
	for dish, rate := range dishesMap {
		fmt.Println("Dish:", dish, "Rate:", rate)
		if rate > maxValue {
			maxKey = dish
			maxValue = rate
		}
	}

	// return the slice
	mostLovedDish[0] = maxKey // return the dish ID
	mostLovedDish[1] = maxValue

	return mostLovedDish
}

func main() {

	//There will be a collection of reviews, each review will have a dish ID and a rating
	reviews := [][]int{{1, 2}, {2, 2}, {3, 4}, {4, 1}, {3, 1}, {4, 5}, {5, 2}}

	//mostLovedDish will return a map with the most loved dish ID and its average rating
	mostLovedDish := mostLovedDish(reviews)
	fmt.Println("The most loved dish ID is:", mostLovedDish[0], " with average rate", mostLovedDish[1]) // Output: {4, 3.0}

}
