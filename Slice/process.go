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
}
