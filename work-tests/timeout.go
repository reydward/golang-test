package main

import (
	"fmt"
	"os"
	"time"
)


func main() {
	graphqlTimeout, _ := time.ParseDuration(os.Getenv("GRAPHQL_TIMEOUTT")+"s")
	if graphqlTimeout == 0 {
		graphqlTimeout = 5 * time.Second
	}
	shortDuration := 6 * time.Second
	fmt.Println(graphqlTimeout)
	fmt.Println(shortDuration)
}
