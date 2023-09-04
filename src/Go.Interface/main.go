package main

import "fmt"

type englishBot struct{}
type spannishBot struct{}

func main() {

}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// func printGreeting(sb spannishBot){
// 	fmt.Println(sb.getGreeting())
// }

func (englishBot) getGreeting() string {
	return "Hi There"
}

func (spannishBot) getGreeting() string {
	return "Holla"
}
