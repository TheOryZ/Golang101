package conditionals

import "fmt"

func Demo1() {
	var account float64 = 1000
	var requestedAmount float64 = 900

	if requestedAmount > account {
		fmt.Println("Your account balance is insufficient")
	} else {
		fmt.Println("Your money is getting ready..")
		account = account - requestedAmount
	}
	fmt.Println("Done! Your account balance :" + fmt.Sprintf("%v", account))
}
