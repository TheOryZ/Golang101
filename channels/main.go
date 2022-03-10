package channels

func EvenNumbers(evenNumbersCn chan int) {
	total := 0
	for i := 0; i <= 10; i += 2 {
		total = total + i
	}
	evenNumbersCn <- total
}

func OddNumbers(oddNumbersCn chan int) {
	total := 0
	for i := 1; i < 10; i += 2 {
		total = total + i
	}
	oddNumbersCn <- total
}

func main() {

}
