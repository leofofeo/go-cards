package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck', which is a slice of strings
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func generateNewDeck() deck {
	newDeck := deck{}
	cardValues := [13]string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}
	suits := [4]string{"Spades", "Clubs", "Diamonds", "Hearts"}
	for _, suit := range suits {
		for _, value := range cardValues {
			card := value + " of " + suit
			newDeck = append(newDeck, card)
		}
	}
	return newDeck
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func dealRandom(d deck, handSize int) deck {
	var cards deck
	// Using rand.Seed here as an alternative to rand.NewSource below
	rand.Seed(time.Now().Unix())
	for i := 0; i < handSize; i++ {
		card := d[rand.Intn(len(d))]
		cards = append(cards, card)
	}
	return cards
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	// Using rand.NewSource here as an alternative to rand.Seed above
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
