package main

import "fmt"

func gen(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	fmt.Println("return statement is called ")
	return out
}

func main() {
	c := make(<-chan int)
	c = gen([]int{2, 3, 4, 5})

	for v := range c {
		fmt.Println(v)
	}


}