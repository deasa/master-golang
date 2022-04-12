package main

import (
	"fmt"
	"time"
	"net/http"
)

func main() {
	websites := []string{
		"http://golang.org",
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://stackoverflow.com",
	}

	c := make(chan string)

	for _, website := range websites {
		go checkWebsite(website, c)
	}

	for l := range c {
		go func (link string)  {
			time.Sleep(5 * time.Second) 
			// having the sleep inside this function literal will make sure the right goroutine is sleeping
			// while still allowing checkWebsite to operate as fast as possible
			// the net effect is that each website will be checked five seconds after it returns
			checkWebsite(link, c)
		}(l)
	}
	// equivalent code
	// for {
	// 	go checkWebsite(<-c, c) // take whatever is put in the channel and send it back into the checkWebsite func to re-check that same website
	// }
}

func checkWebsite(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Printf("Website was down: %s\n", url)
		c <- url
	}

	fmt.Printf("Website was up: %s\n", url)
	c <- url
}
