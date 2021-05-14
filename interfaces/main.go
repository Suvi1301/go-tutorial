package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// NOTE: we dont have a receiever value (only type) because we never make use of it in the receiver function.
func (englishBot) getGreeting() string {
	// Very custom logic for generating an english greeting
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	// Very custom logic for generating a spanish greeting
	return "Hola!"
}
