package main

import "fmt"

type contectInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contectInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(name string) {
	(*p).firstName = name
}

func main() {
	//alex := person{"Alex", "Anderson"}
	//alex := person{firstName: "Alex", lastName: "Anderson"}

	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	jim := person{
		firstName: "Jim",
		lastName:  "Pery",
		contact: contectInfo{
			email:   "jim@gmail.com",
			zipCode: 9400,
		},
	}

	//jim.print()
	jimpointer := &jim
	jimpointer.updateName("mario")

	jim.print()
}
