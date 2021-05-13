package main

import "fmt"

func main() {
	cards := newDeck()

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
}
