package main

import "goModuls/restful"

// "fmt"
// "goModuls/closures"
// "goModuls/recursion"
// functions "goModuls/Functions"
// "goModuls/Pointers"
// slice "goModuls/Slice"
// "goModuls/maps"
// "goModuls/ranges"
//conditionals "goModuls/Conditionals"
//loops "goModuls/Loops"
//arrays "goModuls/Arrays"
// "fmt"
// deferstatement "goModuls/DeferStatement"
// "goModuls/channels"
// errorhandling "goModuls/errorHandling"
// "goModuls/interfacess"
// stringfunctions "goModuls/stringFunctions"

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

	// nextINt := closures.IntSeq()

	// fmt.Println(nextINt())
	// fmt.Println(nextINt())
	// fmt.Println(nextINt())

	// newInts := closures.IntSeq()
	// fmt.Println(newInts())
	// fmt.Println("***********************")
	// fmt.Println(recursion.Fact(7))

	// typesandstructs.Process2()
	// go goroutines.EvenNumbers() //async
	// go goroutines.OddNumbers()  //async

	// evenNumbersCn := make(chan int)
	// oddNumbersCn := make(chan int)
	// go channels.EvenNumbers(evenNumbersCn)
	// go channels.OddNumbers(oddNumbersCn)

	// //time.Sleep(4 * time.Second)
	// evenNumbersTotal, oddevenNumbersTotal := <-evenNumbersCn, <-oddNumbersCn
	// result := evenNumbersTotal * oddevenNumbersTotal
	// fmt.Println("Even Numbers : ", evenNumbersTotal)
	// fmt.Println("Odd Numbers : ", oddevenNumbersTotal)
	// fmt.Println("Result : ", result)
	// fmt.Println("Main is done!!")

	// //interfaces..
	// r := interfacess.Rect{Width: 3, Height: 4}
	// c := interfacess.Circle{Radius: 5}

	// interfacess.Measure(r)
	// interfacess.Measure(c)

	// //Defer Statement..
	// deferstatement.B()
	// deferstatement.Demo()

	// //Error Handling
	// errorhandling.Demo()

	// message, err := errorhandling.GuessAgain(158)
	// fmt.Println(message, "ERROR :", err)

	// //stringfunctions
	// stringfunctions.Demo()

	//restful.Process()
	//restful.Process2()
	restful.ContextExample()
}
