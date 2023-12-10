package main

import "fmt"

// below line fails
// hey := 20
// because Variables can be initialized outside of a function, but cannot be assigned a variable.

func main() {
	// Arrays are pass by value, Slices are pass by reference
	// Below 2 lines of code are EQUIVALENT
	// second one relies on Compiler to know the type of variable
	// type inference to a degree

	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card = "Five of Diamonds"

	// card := newCard()
	// cards := deck{newCard(), newCard()}

	// // adding new element to slice
	// // append returns a new slice, doesn't change the current one
	// cards = append(cards, "Six of Spades")
	cards := newDeck()

	fakeDeck := standardDeck()
	// Iterating over the slice of "cards"
	// cards.print()

	// fakeDeck.print()
	fakeDeck.findSuit()
	// fmt.Println(cards)
	cards.print()
	hand, remainingDeck := deal(cards, 5)
	fmt.Println("Hand:")
	hand.print()
	fmt.Println("Remaining of deck: ")
	remainingDeck.print()
}

// Function that returns a string
func newCard() string {
	return "Five of Diamonds"
}
