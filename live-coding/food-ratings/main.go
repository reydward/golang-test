package main

import (
	"fmt"
	"strconv"
)

func solution(n int, ratings [][]int) int {
	fmt.Println(strconv.Itoa(n))

	dishRatings := make(map[int]struct {
		totalRating int
		count       int
	})

	for _, rating := range ratings {
		dishID := rating[0]
		dishRating := rating[1]

		dishRatings[dishID] = struct {
			totalRating int
			count       int
		}{
			totalRating: dishRatings[dishID].totalRating + dishRating,
			count:       dishRatings[dishID].count + 1,
		}
	}

	highestAvgRating := 0.0
	mostLovedDish := -1

	for dishID, ratingInfo := range dishRatings {
		avgRating := float64(ratingInfo.totalRating) / float64(ratingInfo.count)

		if avgRating > highestAvgRating || (avgRating == highestAvgRating && dishID < mostLovedDish) {
			highestAvgRating = avgRating
			mostLovedDish = dishID
		}
	}

	return mostLovedDish
}

func main() {
	// Example usage
	n := 4
	//ratings := [][]int{{512, 3}, {123, 3}, {123, 4}, {987, 4}}
	ratings := [][]int{{987654423, 4}, {987654220, 5}, {987654202, 4}, {987654250, 1}, {987654419, 5}}

	mostLovedDish := solution(n, ratings)
	fmt.Println("The most loved dish ID is:", mostLovedDish) // Output: 123
}

/*
func main() {
	var N int
	fmt.Scanln(&N)
	ratings := make([][]int, N)
	for i_ratings := 0; i_ratings < N; i_ratings++ {
		ratings[i_ratings] = make([]int, 2)
		for j_ratings := 0; j_ratings < 2; j_ratings++ {
			fmt.Scan(&ratings[i_ratings][j_ratings])
		}
	}

	var out_ int = solution(N, ratings)
	fmt.Print(out_)
}
*/
