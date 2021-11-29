package main

/*
Statement
Create a Stack structure that means a LIFO (Last In First Out).
Note: The func main should be provided and Just need to create the Stack struct and Methods to Push, Pop and Peek. IsEmpty is optional.
To test it, push two or three strings and then pop the last one and finally, push another one.
 */

type Stack struct {
	nodes []string
	count int
}

func NewStack() (stack *Stack){
	stack = &Stack{
	}
	return stack
}

func (stack *Stack) Push(element string) {
	stack.nodes = append(stack.nodes, element)
}

func (stack *Stack) Peek() (topElement string){
	return stack.nodes[len(stack.nodes)-1]
}

func (stack *Stack) Pop() (topElement string){
	removedElement := stack.nodes[(len(stack.nodes)-1)]
	stack.nodes = stack.nodes[:len(stack.nodes)-1]
	return removedElement
}
/*
func main() {
	s := NewStack()
	s.Push("Apple")
	fmt.Printf("Current Stack: %+v \n", *s) // Current Stack: {nodes:[Apple] count:1}
	s.Push("Melon")
	fmt.Printf("Current Stack: %+v \n", *s) // Current Stack: {nodes:[Apple Melon] count:2}
	s.Push("Mangoes")
	fmt.Printf("Current Stack: %+v \n", *s) // Current Stack: {nodes:[Apple Melon Mangoes] count:3}
	fmt.Printf("Show top element: (%s)\n", s.Peek())  // Show top element: (Mangoes)
	fmt.Printf("Current Stack: %+v \n", *s) // Current Stack: {nodes:[Apple Melon Mangoes] count:3}
	fmt.Printf("Extract top element: (%s)\n",s.Pop()) // Extract top element: (Mangoes)
	fmt.Printf("Current Stack: %+v \n", *s) // Current Stack: {nodes:[Apple Melon] count:2}
	s.Push("Pineapple")
	fmt.Printf("Current Stack: %+v \n", *s) // Current Stack: {nodes:[Apple Melon Pineapple] count:3}
}*/