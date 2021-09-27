package ranges

import "fmt"

func Process() {
	//for-Range..
	cities := []string{"Sivas", "Paris", "Berlin"}

	for i, city := range cities {
		fmt.Print(i, " ")
		fmt.Println(city)
	}
	//if u don't want use i (index)..
	for _, city := range cities {
		fmt.Println(city)
	}

	//k : key and v : value.. First item always be key.
	dictionary := map[string]string{"book": "kitap", "table": "masa", "pencil": "kalem"}

	for k, v := range dictionary {
		fmt.Print(k, " ")
		fmt.Println(v)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range kvs {
		fmt.Println("key:", k)
	}
}
