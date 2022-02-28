package loops

import "fmt"

func Demo1() {

	var myString string = "Hello world!"

	for i := 0; i < 5; i++ {
		fmt.Println(myString)
	}
	j := 1
	for j <= 2 {
		fmt.Println("Good bye Real world!!")
		j++
	}

	//for but like while
	sum := 1
	for sum <= 10 {
		sum += sum
		fmt.Println(sum)
	}
}
