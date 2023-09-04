package main

import (
	"fmt"
)

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
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	//Adding a new value into an existing map
	colors["white"] = "#ffffff"

	//modify a value in map
	colors["white"] = "#0ffff"

	delete(colors, "white")
	fmt.Println((colors))

	//Iterating over maps

	printMap(colors)

}

func updateSlice(s []string) {
	s[0] = "Bye"
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color + " :" + hex)
	}
}
