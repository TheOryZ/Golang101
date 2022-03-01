package slice

import "fmt"

func Process() {
	names := make([]string, 3)
	names[0] = "John"
	names[1] = "Doe"
	names[2] = "Mark"
	names = append(names, "Tommy")

	fmt.Println(names)
}

func Process2() {
	cities := []string{"Sivas", "Paris", "Berlin"}
	citiesCopy := make([]string, len(cities))
	fmt.Println(citiesCopy)
	cities = append(cities, "Tokyo")
	copy(citiesCopy, cities)
	fmt.Println(cities)
	fmt.Println(citiesCopy)

	fmt.Println(cities[1:3])
	fmt.Println(cities[1:])
	fmt.Println(cities[:2])

	myArray1 := [...]int{1, 2, 3}
	//you can not add any item. If you use ... for length then u should set values to first line
	fmt.Println(myArray1)

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	myArray2 := [...]int{45, 23, 43}
	mySlice1 := myArray2[:]
	fmt.Println(mySlice1)
	mySlice1[0] = 11
	fmt.Println(mySlice1)
	fmt.Println(myArray2)
}
