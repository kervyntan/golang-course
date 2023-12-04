package main

import "fmt"

// below line fails
// hey := 20
// because Variables can be initialized outside of a function, but cannot be assigned a variable.

func main() {
	// Below 2 lines of code are EQUIVALENT
	// second one relies on Compiler to know the type of variable
	// type inference to a degree

	// var card string = "Ace of Spades"
	card := "Ace of Spades"
	card = "Five of Diamonds"

	fmt.Println(card)
}
