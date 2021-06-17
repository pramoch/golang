package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 3 {
		t.Errorf("Expected deck length of 3, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of 'Ace of Spades', but got '%v'", d[0])
	}
}
