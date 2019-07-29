package main

import "testing"

func TestGenerateNewDeck(t *testing.T) {
	d := generateNewDeck()
	if len(d) != 52 {
		t.Error("Expected deck length of 52, but got", len(d))
	}
}
