package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var cards []string
	deck := generateNewDeck()

	numberOfCards := 9
	fmt.Println("Number of cards: ", numberOfCards)
	for i := 0; i < numberOfCards; i++ {
		cards = append(cards, deck[i])
	}
	fmt.Println("Your cards are: ")
	fmt.Println(cards)
}

func getNewCard(deck []string) string {
	rand.Seed(time.Now().Unix())
	card := deck[rand.Intn(len(deck))]
	return card
}

func generateNewDeck() []string {
	var deck []string
	cardValues := [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	suits := [4]string{"Spades", "Clubs", "Diamonds", "Hearts"}
	for _, face := range cardValues {
		for _, suit := range suits {
			card := face + " of " + suit
			deck = append(deck, card)
		}
	}
	return deck
}
