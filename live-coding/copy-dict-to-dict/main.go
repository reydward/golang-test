package main

import "fmt"

/**
copy a dict to another dict. Key and value can be of any kind of type
a[1] = "1"
a[false]= "false"
a["somestr"] = str
*/

func main() {
	fmt.Println("Hello World!")

	a := make(map[interface{}]interface{})
	a[1] = "1"
	a[false] = "false"
	a["somestr"] = "str"

	b := make(map[interface{}]interface{})

	for index, value := range a {
		b[index] = value
	}

	fmt.Println(a)
	fmt.Println(b)

}
