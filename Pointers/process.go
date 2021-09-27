package Pointers

import (
	"fmt"
)

func Process() {
	var myString string = "Green"

	//log.Println("myString is set to", myString)
	fmt.Println("myString is set to", myString)
	changeUsingPointer(&myString)
	fmt.Println("after func call myString is set to", myString)
	fmt.Println("myString is point", &myString)
}

func changeUsingPointer(s *string) {
	fmt.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}
