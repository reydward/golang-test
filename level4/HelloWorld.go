package main

/*
Statement
Create a goroutine that will execute an anonymous function to print “Hello World” and in the main routine print “main function”

Topics to Practice:
goroutine, function, common Issue

var wg sync.WaitGroup

func main(){
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())

	wg.Add(1)
	go func() {
		fmt.Println("Hello World")
		wg.Done()
	}()

	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	time.Sleep(time.Duration(2))
	fmt.Println("Main function")

	wg.Wait()
}
*/