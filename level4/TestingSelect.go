package main

import "fmt"

func main()  {
	pair := make(chan int)
	odd := make(chan int)
	exit := make(chan int)

	go send(pair, odd, exit)

	receive(pair, odd, exit)

	fmt.Println("Finished")
}

func send(pair, odd, exit chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			pair <- i
		} else {
			odd <- i
		}
	}
 	exit <- 0
}

func receive(pair, odd, exit <-chan int) {
	for{
		select {
		case v := <- pair:
			fmt.Println("Pair channel: ", v)
		case v := <- odd:
			fmt.Println("Odd channel: ", v)
		case v := <- exit:
			fmt.Println("Exit channel: ", v)
			return
		}
	}
}
