package main

import (
	"fmt"
	"strings"
)

// Create a new type of "deck"
// which is a slice of strings (borrows behaviour from a slice of string)
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func standardDeck() deck {
	// stdDeck is an empty slice with elements of type string
	stdDeck := deck{}

	cardOrder := []string{"First", "Second", "Third", "Fourth"}
	cardValues := []string{"Hades", "Hercules", "Romans"}

	for _, order := range cardOrder {
		for _, value := range cardValues {
			stdDeck = append(stdDeck, order+" of "+value)
		}
	}

	return stdDeck
}

func (d deck) findSuit() {
	for _, card := range d {
		// from index of "of", add 2 to start on the 1st character of the suit
		indexToStart := strings.Index(card, "of") + 2
		fmt.Println(card[indexToStart:])
	}
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
