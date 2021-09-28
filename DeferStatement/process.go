package deferstatement

import "fmt"

func A() {
	fmt.Println("function A worked")
}
func B() {
	defer A() //defer runs at the very end of the code it's in, no matter where it is.
	defer C() //last in first out!!
	fmt.Println("function B worked")
}
func C() {
	fmt.Println("function C worked")
}

/*
	When we encounter an error in the function, the functions we call with the word defer will still work.
	We can use the defer in logging algorithms.
*/
