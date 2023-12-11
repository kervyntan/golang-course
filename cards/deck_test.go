package main

import "testing"

/*
Code to make sure that a deck is created with a specified
number of cards
*/
func TestNewDeck(t *testing.T) {
	d := newDeck()

	// 4 types of values, 4 suits = 4 * 4 = 16 cards
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
}
