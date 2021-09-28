package errorhandling

import (
	"fmt"
	"os"
)

func Demo() {
	file, err := os.Open("errorhandling/demo1.txt") // if file is exist, err = nil

	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("File is not exist! PATH : ", pErr.Path)
			return
		} else {
			fmt.Println("File is not exist!")
			return
		}
	} else {
		fmt.Println(file.Name())
	}
}
