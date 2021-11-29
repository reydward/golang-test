package main

import "fmt"

/*
Create a Trumpeter struct and a Violinist struct both struts should have a Name attribute and a Greetings() function to present themself.
Then create a MusicalPlayer interface.
In the main function, create a slice with two or more musical players and for each call the Greetings() function.
Topics to Practice:
Type Interfaces, struct, method, function, slice, for each loop
 */

type Trumpeter struct {
	name string
}

type Violinist struct {
	name string
}

func (t Trumpeter) Greetings() {
	fmt.Println("I am trumpeter")
}

func (v Violinist) Greetings() {
	fmt.Println("I am violinist")
}

type MusicalPlayer interface {
	Greetings()
}
