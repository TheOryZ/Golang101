package main

import "fmt"

func main() {
	str1 := "John Doe"
	str2 := "Kadir Turan"
	str3 := "TheOryZ"

	stringLength, _ := fmt.Println(str1, str2, str3)
	fmt.Println(stringLength)

	aNumber := 58
	isTrue := true

	fmt.Printf("Value of aNumber: %v\n", aNumber)
	fmt.Printf("Value of isTrue: %v\n", isTrue)
	fmt.Printf("Value of aNumber as float: %.2f\n", float64(aNumber))
	fmt.Printf("Data types: %T, %T, %T, %T, and %T\n", str1, str2, str3, aNumber, isTrue)
	mystring := fmt.Sprintf("Data types as var: %T, %T, %T, %T, and %T", str1, str2, str3, aNumber, isTrue)
	fmt.Print(mystring)
}
