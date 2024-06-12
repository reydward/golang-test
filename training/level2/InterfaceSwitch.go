package main

import "fmt"

/*
Interface Switch
Statement
Create a function that will receive an interface and it will print the interface value with a specific message based on the type.

Topics to Practice:
Interfaces, type assertion, switch
*/

type type1 struct {
	attribute1 string
}

type type2 struct {
	attribute2 string
}

func (t type1) getValue()  {
	fmt.Println("value from type1")
}

func (t type2) getValue()  {
	fmt.Println("value from type2")
}

type interface1 interface {
	getValue()
}

func printMessage(interface1 interface1) {
	switch interface1.(type) {
	case type1:
		fmt.Println("type1: " + interface1.(type1).attribute1)
	case type2:
		fmt.Println("type2: " + interface1.(type2).attribute2)
	}
}

