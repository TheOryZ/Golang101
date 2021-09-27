package typesandstructs

import "fmt"

func Process2() {
	c := customer{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}
	c.save()
	c.update()
}

func (c customer) save() {
	fmt.Println("Customer Added.", c.FirstName, c.LastName)
}

func (c customer) update() {
	fmt.Println("Customer Updated.", c.FirstName, c.LastName)
}

type customer struct {
	FirstName string
	LastName  string
	Age       int
}
