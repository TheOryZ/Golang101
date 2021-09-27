package maps

import "fmt"

func Process() {
	//key-value
	dictionary := make(map[string]string)
	dictionary["table"] = "masa"
	dictionary["book"] = "kitap"
	dictionary["pencil"] = "kalem"

	fmt.Println(dictionary["book"])
	fmt.Println("Dictionary items count : ", len(dictionary))
	delete(dictionary, "pencil")
	fmt.Println("Dictionary : ", dictionary)

	//searching..
	value, isThere := dictionary["pencil"]
	fmt.Println(value)
	fmt.Println("True or False : ", isThere)

	//without make..
	dictionary2 := map[string]string{"glass": "bardak", "microphone": "mikrofon"}
	fmt.Println(dictionary2)
}
