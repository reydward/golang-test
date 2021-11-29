package main

import (
	"fmt"
	"sync"
)

func main()  {
	var wg sync.WaitGroup
	increment := 0
	gs := 100
	wg.Add(gs)
	var mutex sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			mutex.Lock()
			v := increment
			v++
			increment = v
			fmt.Println(increment)
			mutex.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Final value:", increment)
}
