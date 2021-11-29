package main

/*
Statement
Create a function that will increase a number and that function will be executed by a goroutine inside a for loop (x1000 times). To avoid race conditioning, implement the sync.Mutex and Lock and Unlock inside the increase() function.
Note: Add a time.Sleep() to be able to see the final n

Topics to Practice:
goroutine, sync.mutex, function, for loop, defer, time pkg

var wg sync.WaitGroup

func increaseNumber(number *int) {
	*number++
	runtime.Gosched()
	wg.Done()
}

func main()  {
	count := 0
	const gr = 100
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go increaseNumber(&count)
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("count: ", count)
}
*/