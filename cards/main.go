package main

import "fmt"

func main() {
	// Below 2 lines of code are EQUIVALENT
	// second one relies on Compiler to know the type of variable

	// var card string = "Ace of Spades"
	card := "Ace of Spades"

	fmt.Println(card)
}
