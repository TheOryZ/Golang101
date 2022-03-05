package functions

import "errors"

func VariadicFunc(numbers ...int) int {
	total := 0
	for i := 0; i < len(numbers); i++ {
		total = total + numbers[i]
	}
	return total
}
func divide(numerator, denominator int) (int, error) {
	if denominator == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, nil
}

//anonim functions
//bir fonksiyon parametre olarak fonksiyon alabiliyor. bununla ilgili bir örnek bul ve eklemesi tamamla..
