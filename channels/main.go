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

	c := make(chan string) // Make a channel of type string.

	for _, link := range links {
		go checkLink(link, c) // launch a new go routine for every blocking func call.
	}

	/*
		This receives the first message from the channel and then main exits.
		This is essentially a blocking call until we get the first message back.
	*/
	// fmt.Println(<-c)

	/*
		We can loop the number of links we know we are calling and wait for that many channel responses.
	*/
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	/*
		Lets try to fetch the links infinitely. Everytime we get a response from one link, we will
		listen to the channel message (which is the link itself), and checkLink again.
	*/
	// for {
	// 	go checkLink(<-c, c)
	// }

	/*
		This is equivalent to the for { } above. Just stylistically better. range c waits for value from channel
		then assigns to l. So steps the for loop everytime channel emits something. Same as above.
	*/
	// for l := range c {
	// 	go checkLink(l, c)
	// }

	// We want to have a delay of 5 seconds before we make the checkLink call to same link again.
	// We use a "function literal" (lambdas in python) with time.Sleep.
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) // Here we pass l to the function literal -> this means its a copy of the original l -> prevents same variable being used in multiple go routines.
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // our "blocking call"
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
