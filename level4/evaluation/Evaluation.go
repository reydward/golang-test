package main

import (
	"fmt"
	"strings"
	"sync"
)

/*
Given a list of strings find all strings that contain a given substring. You should split the list into chunks and process them in parallel.
Each goroutine should write the response to an output channel which will be consumed in the end to consolidate the response.

strs := []string{"abc", "bcd", "efg", "aabcr", "acb", "ggg", "hjuklbc"}
substr := "bc"

strs := []string{"abc", "bcd", "aabcr", "hjuklbc"}

 */
const substr = "bc"
var wg sync.WaitGroup

func main()  {
	slice := []string{"abc", "bcd", "efg", "aabcr", "acb", "ggg", "hjuklbc", "bcbc"}
	finalResult := make(chan string)

	chunkSize := 3

	//Split the list
	var end int
	for start := 0; start < len(slice); start += chunkSize{
		end += chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		fmt.Println("Chunk: ", slice[start:end])
		//Process in parallel
		//Write the response to an output channel
		wg.Add(1)
		go processSlice(slice[start:end], finalResult)
	}

	//Consuming to consolidate response
	go readFinalResult(finalResult)
	wg.Wait()
	fmt.Println("This is the end!")
}

func numRoutines(sliceSize int, chunkSize int) int {
	return (sliceSize / chunkSize) + (sliceSize % chunkSize)
}

func processSlice(slice []string, finalResult chan<- string) {
//	fmt.Println("processSlice.slice: ", slice)
	for _, v := range slice {
		if strings.Contains(v, substr) {
			finalResult <- v
		}
	}
	wg.Done()
}

func readFinalResult(finalResult <-chan string) {
	fmt.Println("Reading final result")
	for{
		select {
		case v := <- finalResult:
			fmt.Println("[", v, "]")
		}
	}
}


