package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck size of 52, but got %v", len(d))
	}

	// Check first card is Ace of Spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	// Check last card is King of Diamonds
	if d[len(d)-1] != "King of Diamonds" {
		t.Errorf("Expected last card to be King of Diamonds, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	testFileName := "_decktesting"

	// clean up any left over files.
	os.Remove(testFileName)
	d := newDeck()
	d.saveToFile(testFileName)

	loadedDeck := newDeckFromFile(testFileName)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in the deck, got %v", len(loadedDeck))
	}

	os.Remove(testFileName)
}
