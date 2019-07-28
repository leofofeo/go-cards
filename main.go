package main

import "fmt"

func main() {
	deck := generateNewDeck()
	fmt.Println(deck)

}

// func getNewCard() string {
// 	deck := generateNewDeck()
// 	fmt.Println(deck)
// }

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
