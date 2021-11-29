package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main()  {
	fmt.Printf("Number of CPUs: %v\n", runtime.NumCPU())
	fmt.Printf("Number of routines: %v\n", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("From go routine 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("From go routine 2")
		wg.Done()
	}()

	fmt.Printf("Number of CPUs: %v\n", runtime.NumCPU())
	fmt.Printf("Number of routines: %v\n", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("Before finish...")

	fmt.Printf("Number of CPUs: %v\n", runtime.NumCPU())
	fmt.Printf("Number of routines: %v\n", runtime.NumGoroutine())
}
