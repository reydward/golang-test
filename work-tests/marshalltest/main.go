package main

import "fmt"

func makeSquares(){
	fmt.Println("hi")
}

const (
	i=7
	j
	k
)

func main()  {
/*a := []string{"a","b","c","d","e"}
a = a[:0]
fmt.Println(a, len(a), cap(a))
*/
	/*a := [...]int{1,2,3,4,5,6,7,8,9}
	s := a[2:4]
	s[0] = 33
	fmt.Println(a[2])
*/
/*	s1 := []int{0,1,2,3}
	s2 := []int{0,1,2,3}
	fmt.Println(s1 == s2)*/

	s2 := []int{1,2,3}
	s3 := []int{4,5,6,7}
	n2 := copy(s2, s3)
	fmt.Printf("%d, %v, %v", n2,s2,s3)
}
