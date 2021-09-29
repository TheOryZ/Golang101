package errorhandling

import "fmt"

type borderException struct {
	parameter int
	message   string
}

func (b *borderException) Error() string {
	return fmt.Sprintf("%d---%s", b.parameter, b.message)
}

func GuessAgain(guess int) (string, error) {
	point := 67
	if guess < 1 || guess > 100 {
		return "", &borderException{guess, "Please enter a number between 1 and 100"}
	} else if guess < point {
		return "", &borderException{guess, "Please enter a larger number."}
	} else if guess > point {
		return "", &borderException{guess, "Please enter a smaller number."}
	} else {
		return "Congratulations! You know.", nil
	}
}
