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
	cards = getNewCards(deck, numberOfCards)
	fmt.Println("Your cards are: ")
	fmt.Println(cards)
}

func getNewCards(deck []string, numberOfCards int) []string {
	var cards []string
	rand.Seed(time.Now().Unix())
	for i := 0; i < numberOfCards; i++ {
		card := deck[rand.Intn(len(deck))]
		cards = append(cards, card)
	}
	return cards
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
