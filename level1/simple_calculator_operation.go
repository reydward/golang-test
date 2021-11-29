/*
Simple Calculator Operations
Statement
Create a calculator struct with a result attribute and with the most common methods:
Add
Subtract
Multiply
Divide


Plus: Handle possible error, Divide function won’t be able to handle second parameter with value zero,
instead function should panic and print an alert

Topics to Practice:
function, return value, panic, defer
*/

package level1

import "fmt"

type Calculator struct {
	X      int
	Y      int
	Result float64
}

func Add(x int, y int) (result Calculator){
	fmt.Println("\n---------------------- Add")
	calculatorResult := Calculator{
		x,
		y,
		float64(x+y),
	}
	return calculatorResult
}

func Substract(x int, y int) (result Calculator){
	fmt.Println("\n---------------------- Substract")
	calculatorResult := Calculator{
		x,
		y,
		float64(x-y),
	}
	return calculatorResult
}

func Multiply(x int, y int) (result Calculator){
	fmt.Println("\n---------------------- Multiply")
	calculatorResult := Calculator{
		x,
		y,
		float64(x*y),
	}
	return calculatorResult
}

func Divide(x int, y int) (result Calculator){
	fmt.Println("\n---------------------- Divide")
	defer func() {
		recovered := recover()
		if recovered != nil {
			fmt.Println("\nYou can not divide by zero!!")
		}
		fmt.Println("\n---------------------- Divide.defer")
	}()

	if y == 0 {
		panic(fmt.Sprintf("Division by zero"))
	}

	calculatorResult := Calculator{
		x,
		y,
		float64(x)/float64(y),
	}
	fmt.Println("\n---------------------- Divide.return")
	return calculatorResult
}