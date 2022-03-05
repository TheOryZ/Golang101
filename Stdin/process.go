package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(len(text))
	fmt.Println(text)
	if len(text) > 3 {
		fmt.Println("Lütfen Sadece harf giriniz")
	} else {
		fmt.Println("Harf girildi.")
	}
}
