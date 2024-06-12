package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	absPath, _ := filepath.Abs("./file.dat")

	file, err := os.Open(absPath)
	//	file, err := os.Open("/Users/eduardreyes/Workspace/golang-test/work-tests/files/data/file.dat")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println(file.Name())

}
