package loops

import "fmt"

func Demo1() {

	var myString string = "Hellor world!"

	for i := 0; i < 5; i++ {
		fmt.Println(myString)
	}
	j := 1
	for j <= 2 {
		fmt.Println("Good bye Real world!!")
		j++
	}

	//fmt.Scanln();
	scanVariables := 0
	fmt.Scanln(&scanVariables)
	fmt.Println(scanVariables)
}
