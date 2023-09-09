package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spannishBot struct{}

func main() {
	tr := triangle{base: 15.0, height: 5.0}
	sq := square{5.3}

	printArea(tr)
	fmt.Println()
	printArea(sq)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi There"
}

func (spannishBot) getGreeting() string {
	return "Holla"
}
