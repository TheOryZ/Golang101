package main

import (
	"fmt"
)

func main() {
	var myString string = "Green"

	//log.Println("myString is set to", myString)
	fmt.Println("myString is set to", myString)
	ChangeUsingPointer(&myString)
	fmt.Println("after func call myString is set to", myString)
	fmt.Println("myString is point", &myString)
}

func ChangeUsingPointer(s *string) {
	fmt.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}
