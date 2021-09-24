package main

import "fmt"

var varMyString string //Kullanmamış olmamıza rağmen hata vermedi. Herhangi bir fonksiyon içerisinde olmadığı için.

func main() {

	fmt.Println("Hello World!")

	var whatToSay string
	whatToSay = "Goodbye, cruel world"
	fmt.Println(whatToSay)

	var i int
	i = 58
	fmt.Println("i is set to ", i)
	whatWasSaid := saySomething()
	whatWasSaidAgain, theOtherThingThatWasSaid := saySomething2()
	fmt.Println("What was said: ", saySomething())
	fmt.Println("What was said again: ", whatWasSaid)
	fmt.Println(whatWasSaidAgain, theOtherThingThatWasSaid)
}

func saySomething() string {
	return "something"
}

func saySomething2() (string, string) {
	return "something", "else"
}
