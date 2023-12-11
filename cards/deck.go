package main

import (
	"fmt"
	"os"
	"strings"
)

/*
- Create a new type of "deck" which is a slice of strings
(borrows behaviour from a slice of string)
*/
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

/*
- Receiver functions can be accessed via the use of a deck type
*/
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

/*
- Return multiple values from a function
- return 2 values, both of TYPE deck
*/
func deal(d deck, handSize int) (deck, deck) {
	// Reference diagram in Video no.24 in the course
	return d[:handSize], d[handSize:]
}

/*
- Process deck into a string, then can turn that into byte slice
separately in another function
*/
func (d deck) toString() string {
	// convert deck to type of string slice
	// join the strings in the slice
	return strings.Join([]string(d), ",")
}

/*
- Writes a file to the directory containing cards in the deck
*/
func (d deck) saveToFile(filename string) error {
	// Convert deck to byte slice
	byteSlice := []byte(d.toString())
	// 0666 grants all permissions to the user
	return os.WriteFile(filename, byteSlice, 0666)
}

/*
- Get a deck of cards from a file
*/
func newDeckFromFile(filename string) deck {
	// ReadFile returns two things
	// if success, err will have a value of nil
	byteSlice, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		// Exit argument: 0 if error-free, other value otherwise
		os.Exit(1)
	} else {
		// Can look at documentation for Split
		s := strings.Split(string(byteSlice), ",")
		return deck(s)
	}
}
