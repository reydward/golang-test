package level1

import "fmt"

func A() {
	i := 0
	defer fmt.Println(i)
	i++
	fmt.Println(i)
	return
}

func B() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

func C() (i int) {
	defer func() { i++ }()
	return 1
}

