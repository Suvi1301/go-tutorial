package main

import (
	"fmt"
	"strings"
)

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}

func notes() {
	// var card string = "Ace of spades"
	card := "Ace of spades" // identical to the line above
	// := should only be used for defining new variables not reassignment.
	fmt.Println(card)

	card1 := newCard()
	fmt.Println(card1)

	// declaring a slice (an array that is flexible)
	cards := []string{newCard(), newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades")

	// Type conversion: string to byte slice
	greeting := "Hi there!"
	fmt.Println([]byte(greeting))

	// Convert a slice of strings into a string.
	convertedString := strings.Join(cards, ",")
	fmt.Println(convertedString)

	// Using the save and read to/from file
	newDeck := newDeck()
	newDeck.saveToFile("my_cards")
	readDeck := newDeckFromFile("my_cards")
	readDeck.print()

	// Testing:
	// - a test file should end with '_test.go'
	// - 'go test' runs all tests in a package
}
