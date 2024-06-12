package main

/*
Statement
Create a channel of int and then create a goroutine to add a value to the channel and then print the channel value in the main function

Topics to Practice:
goroutine, channel, function
func main()  {

	myChannel := make(chan int)

	go func() {
		myChannel <- 42
	}()

	fmt.Println(<- myChannel)
}

 */