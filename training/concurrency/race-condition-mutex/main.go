package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Número de CPUs:", runtime.NumCPU())
	fmt.Println("Número de Goroutines:", runtime.NumGoroutine())

	contador := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := contador
			v++
			runtime.Gosched()
			contador = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Número de Goroutinas:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Cuenta:", contador)
}
