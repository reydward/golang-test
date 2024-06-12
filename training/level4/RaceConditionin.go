package main
/*
import (
	"fmt"
	"runtime"
	"sync"
)
func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	count := 0
	const gr = 100
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			mutex.Lock()
			v := count
			v++
			runtime.Gosched()
			count = v
			mutex.Unlock()
			wg.Done()
		}()
		fmt.Println("--- Goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Count: ", count)
}
*/