package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {

	person1 := Person{"John", 30, "email1@email.com"}
	person2 := Person{"John", 30, "email1@email.com"}
	person3 := Person{"John", 30, "email1@email.com"}

	persons := []Person{person1, person2, person3}

	for _, person := range persons {
		fmt.Println(person)
	}

}
