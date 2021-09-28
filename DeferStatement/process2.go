package deferstatement

import "fmt"

type product struct {
	productName string
	unitPrice   float64
}

func (p product) save() {
	fmt.Println("Product is saved : ", p.productName)
	defer log()
}

func Demo() {
	p := product{productName: "Laptop", unitPrice: 5000.0}
	defer p.save() //It is necessary to pay attention to the point where Defer is defined.
	p = product{productName: "Maouse", unitPrice: 40.0}
	fmt.Println("Process is done!")
}

func log() {
	fmt.Println("Logging...")
}
