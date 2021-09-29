package errorhandling

import (
	"fmt"
)

func Process2() {
	message, err := Guess(158)
	fmt.Println(message, "ERROR :", err)
}

func Guess(guess int) (string, error) {
	point := 58
	if guess < 1 || guess > 100 {
		return "Please enter a number between 1 and 100", nil //errors.New("Please enter a number between 1 and 100")
	} else if guess < point {
		return "Please enter a larger number.", nil
	} else if guess > point {
		return "Please enter a smaller number", nil
	} else {
		return "Congratulations! You know.", nil
	}
}
