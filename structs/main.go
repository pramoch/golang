package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// #1
	// alex := person{"Alex", "Anderson"}

	// #2
	// michael := person{firstName: "Michael", lastName: "Anderson"}

	// #3
	// var david person
	// david.firstName = "David"
	// david.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Println(michael)
	// fmt.Printf("%+v", david)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateNameByValue("Jimmy")
	jim.print()

	jimPointer := &jim
	jimPointer.updateNameByReference("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}

func (p person) updateNameByValue(newFirstName string) {
	p.firstName = newFirstName
}

func (pointerToPersion *person) updateNameByReference(newFirstName string) {
	(*pointerToPersion).firstName = newFirstName
}
