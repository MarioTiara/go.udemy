package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoveflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}
	// for {
	// 	go checkLink(<-c, c)
	// }

	for l := range c {
		time.Sleep(5 * time.Second)
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- "Might be down, I think"
	}

	fmt.Println(link, "is up")
	c <- "Yup is ts aup"
}
