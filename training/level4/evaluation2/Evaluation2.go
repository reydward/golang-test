package main
/*
Given a list of strings find all strings that contain a given substring. You should split the list into chunks and process them in parallel.
Each goroutine should write the response to an output channel which will be consumed in the end to consolidate the response.

strs := []string{"abc", "bcd", "efg", "aabcr", "acb", "ggg", "hjuklbc"}
substr := "bc"

strs := []string{"abc", "bcd", "aabcr", "hjuklbc"}
*/
import (
	"fmt"
	"sync"
)

const substr = "bc"

func main()  {
	slice := []string{"abc", "bcd", "efg", "aabcr", "acb", "ggg", "hjuklbc"}
	const chunkSize = 3
	//finalResult := make(chan string)

	var wg sync.WaitGroup
	wg.Add(3)

	//Split the list
	var end int
	var subslice []string
	for start := 0; start < len(slice); start += chunkSize{
		end += chunkSize
		if end > len(slice) {
			end = len(slice)
		}
/*		for _, v := range slice[start:end] {
			if strings.Contains(v, substr) {
				fmt.Println(v)
			}
		}
*/
		//Process in parallel
		subslice = slice[start:end]
		fmt.Println("Chunk1: ", subslice)
		//Write the response to an output channel
		go func() {
			fmt.Println("Go function", subslice)
			for _, v := range subslice {
				fmt.Println(v)
/*				if strings.Contains(v, substr) {
					finalResult <- v
				}*/
			}
			wg.Done()
		}()
		fmt.Println("Chunk2: ", subslice)

	}

	wg.Wait()
}
