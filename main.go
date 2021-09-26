package main

import (
	//arrays "goModuls/Arrays"
	"fmt"
	functions "goModuls/Functions"
	slice "goModuls/Slice"
	//conditionals "goModuls/Conditionals"
	//loops "goModuls/Loops"
)

func main() {
	//conditionals.Demo1()
	//loops.Demo1()
	//arrays.Demo1()
	slice.Process()
	slice.Process2()
	array_ := []int{4, 8, 15, 16, 23, 42}
	fmt.Println(functions.VariadicFunc(1, 2, 3, 4, 5))
	fmt.Println(functions.VariadicFunc(array_...))
	fmt.Println(functions.VariadicFunc(1))
	fmt.Println(functions.VariadicFunc())
}