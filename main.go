package main

import (
	"fmt"
)

func main() {
	// deck := generateNewDeck()
	// deck.print()
	// numberOfCards := 5
	// hand, remainingCards := deal(deck, numberOfCards)
	// fmt.Println("Your cards are: ")
	// hand.print()
	// fmt.Println("Remaining cards are: ")
	// remainingCards.print()
	cards := generateNewDeck()
	fmt.Println(cards.toString())

}
