package main

import "fmt"

// A deferred function’s arguments are evaluated when the defer statement is evaluated
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

// Deferred function calls are executed in Last In First Out order after the surrounding function returns.
func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

// Deferred functions may read and assign to the returning function’s named return values.
// In this example, a deferred function increments the return value i after the surrounding function returns.
// Thus, this function returns 2:
func c() (i int) {
	defer func() { i++ }()
	return 1
}
func main() {
	a()
	b()
	c := c()
	fmt.Println(c)
}
