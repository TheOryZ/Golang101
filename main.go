package main

import (
	"fmt"
	"goModuls/closures"
	"goModuls/recursion"
	// functions "goModuls/Functions"
	// "goModuls/Pointers"
	// slice "goModuls/Slice"
	// "goModuls/maps"
	// "goModuls/ranges"
	//conditionals "goModuls/Conditionals"
	//loops "goModuls/Loops"
	//arrays "goModuls/Arrays"
)

func main() {
	//conditionals.Demo1()
	//loops.Demo1()
	//arrays.Demo1()
	// slice.Process()
	// slice.Process2()
	// array_ := []int{4, 8, 15, 16, 23, 42}
	// fmt.Println(functions.VariadicFunc(1, 2, 3, 4, 5))
	// fmt.Println(functions.VariadicFunc(array_...))
	// fmt.Println(functions.VariadicFunc(1))
	// fmt.Println(functions.VariadicFunc())
	// Pointers.Process()
	// maps.Process()
	// ranges.Process()
	nextINt := closures.IntSeq()

	fmt.Println(nextINt())
	fmt.Println(nextINt())
	fmt.Println(nextINt())

	newInts := closures.IntSeq()
	fmt.Println(newInts())
	fmt.Println("***********************")
	fmt.Println(recursion.Fact(7))

}
