package interfacess

import "fmt"

func verify(i interface{}) {
	number, ok := i.(int)
	fmt.Println(number, ok)
}

func Control() {
	var checkNumber interface{} = "John Doe"
	//var checkNumber int = 5
	verify(checkNumber)

	var checkNumber2 int = 5
	verify(checkNumber2)
}
