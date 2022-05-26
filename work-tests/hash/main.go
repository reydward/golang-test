package main

import (
	"fmt"
	"hash/fnv"
)

func HashVin(vin string) string {
	h := fnv.New32a()
	h.Write([]byte(vin))
	return fmt.Sprint(h.Sum32())
}

func main()  {
	hashVin := HashVin("2C3CCABGXKH617963")
	fmt.Println(hashVin)
}
