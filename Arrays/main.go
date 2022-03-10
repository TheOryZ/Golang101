package arrays

import "fmt"

func Demo1() {
	var numbers [5]int
	numbers[2] = 58
	fmt.Println(numbers)

	//var cities [5]string

	numbers2 := [5]int{5, 49, 78, 15, 35}

	for i := 0; i < len(numbers2); i++ {
		fmt.Println(numbers2[i])
	}

	//multidimensional arrays
	var multiNumbers [2][3]int
	multiNumbers[0][0] = 1
	multiNumbers[0][1] = 2
	multiNumbers[0][2] = 3
	multiNumbers[1][0] = 2
	multiNumbers[1][1] = 3
	multiNumbers[1][2] = 4

	fmt.Println(multiNumbers)
}
