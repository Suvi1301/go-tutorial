package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// (d deck) syntax tells us this is a receiver. So any variable of type deck can access this function.
// deck: the type we want to assign this func to.
// d: instance of the deck variable we are working with. So calling cards.print() essentially passes the reference to cards as 'd'.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
