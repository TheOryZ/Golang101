package goroutines

import "fmt"

func EvenNumbers() {
	for i := 0; i <= 10; i += 2 {
		fmt.Println("Even Number : ", i)
	}
}

func OddNumbers() {
	for i := 1; i < 10; i += 2 {
		fmt.Println("Odd Number : ", i)
	}
}
