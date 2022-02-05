package main

import (
	"fmt"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) printcards() {

	for i, deck := range d {
		fmt.Println(i, deck)
	}
}

func deal(d deck, hand int) (deck, deck) {

	return d[:hand], d[hand:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
