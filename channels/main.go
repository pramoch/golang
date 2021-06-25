package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://stackoverflow.com",
		"https://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// #1 - Loop 1 time
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// #2 - Loop forever
	// for {
	// 	go checkLink(<-c, c)
	// }

	// #3 - Alternative syntax of loop forever
	for l := range c {
		// Function literal (like anonymous function in Javascript)
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "migh be down")
		c <- link
		return
	}

	fmt.Println(link, "is up")
	c <- link
}
