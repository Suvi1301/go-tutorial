package main

import (
	"fmt"
	"strings"
)

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
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
}
