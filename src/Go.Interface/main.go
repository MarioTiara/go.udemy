package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spannishBot struct{}

func main() {
	eb := englishBot{}
	sb := spannishBot{}

	printGreeting(eb)
	printGreeting(sb)
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
