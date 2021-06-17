package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// #1
	alex := person{"Alex", "Anderson"}

	// #2
	michael := person{firstName: "Michael", lastName: "Anderson"}

	// #3
	var david person
	david.firstName = "David"
	david.lastName = "Anderson"

	fmt.Println(alex)
	fmt.Println(michael)
	fmt.Printf("%+v", david)
}
