package main

import (
	"fmt"
	"os"
)

func main()  {
	fmt.Println("--------------------- Path")
	fmt.Println(os.Getenv("Path"))
	fmt.Println("--------------------- SHOPIFY_QUEUE_URL")
	fmt.Println(os.Getenv("SHOPIFY_QUEUE_URL"))
	fmt.Println("--------------------- AWS_REGION")
	fmt.Println(os.Getenv("AWS_REGION"))
}
