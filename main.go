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
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, value := range links {
		go checkLink(value, c)
	}

	for l := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " Might be down!")
		channel <- link
		return
	}
	fmt.Println(link + " is Up!")
	channel <- link
}
