package main

import "fmt"

type contectInfo struct {
	email   []string
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

func (p *person) UpdateFirstEmail(email string) {
	if len(p.contact.email) > 0 {
		p.contact.email[0] = email
	}
}

func main() {
	email := []string{"mario@gmail.com", "mario@microsoft.com"}

	mario := person{
		firstName: "mario",
		lastName:  "pratama",
		contact: contectInfo{
			email:   email,
			zipCode: 3232,
		},
	}

	mario.UpdateFirstEmail("putra@gmail.com")

	fmt.Println(mario)

}

func updateSlice(s []string) {
	s[0] = "Bye"
}
