package main

import (
	"context"
	"fmt"
	"time"
)
const timeout = 3 * time.Second

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	errors := make(chan int, 1)
	errors <- 333

	select {
	case err := <-errors:
		fmt.Println(err)
	case <-ctx.Done():
		fmt.Println("------------------------> timeout")
	}
}

