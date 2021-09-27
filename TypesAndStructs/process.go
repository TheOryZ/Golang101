package typesandstructs

import (
	"fmt"
	"time"
)

//var s = "seven"

func Process() {
	s := "eight"
	/*Aynı isimde bir başka değişkeni, alt fonksiyonlarda tanımlayabilirsin.
	Daha içteki işlem bloğundaki tanımlama baskın gelecektir.
	Tanımlama yaparken dikkatli isim ver!!*/
	var s2 = "six"
	fmt.Println("s is", s)
	fmt.Println("s2 is", s2)

	saySomething(("xxx"))

	/* ---TYPES--- */
	user := user{
		FirstName: "John",
		LastName:  "Doe",
	}

	fmt.Println(user.FirstName, user.LastName, user.BirthDate, user.Age)

	product := product{
		name:      "Personel computer",
		unitPrice: 15000,
		brand:     "DELL",
	}

	fmt.Println(product)
}

/*Erişim belirleyici keyleri (public, private vs.) yoktur.
Eğer obje büyük harf ile başlarsa Public
Eğer obje küçük harf ile başlarsa Private
*/
func saySomething(s string) (string, string) {
	fmt.Println("s from the saySomethings func is", s)
	return s, "World"
}

type user struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

type product struct {
	name      string
	unitPrice float64
	brand     string
}
