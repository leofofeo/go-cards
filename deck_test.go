package main

import (
	"os"
	"testing"
)

func TestGenerateNewDeck(t *testing.T) {
	d := generateNewDeck()
	if len(d) != 52 {
		t.Error("Expected deck length of 52, but got", len(d))
	}

	if d[0] != "Two of Spades" {
		t.Error("Expected 'Two of Spades' as first card, but received", d[0])
	}

	if d[len(d)-1] != "Ace of Hearts" {
		t.Error("Expected 'Ace of Hearts' as last card, but received", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := generateNewDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Error("Expected 52 cards in deck, received", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
