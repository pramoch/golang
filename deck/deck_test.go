package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 3 {
		t.Errorf("Expected deck length of 3, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of 'Ace of Spades', but got '%v'", d[0])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 3 {
		t.Errorf("Expected deck length of 3, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
