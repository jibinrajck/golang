package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
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

func (d deck) saveToFile(sFilename string) error {
	return ioutil.WriteFile(sFilename, []byte(d.toString()), 0666)
}

func newDeckFromFile(sFileName string) deck {
	ba, error := ioutil.ReadFile(sFileName)

	if error != nil {
		fmt.Println("Error: ", error)
		os.Exit(1)
	}

	return deck(strings.Split(string(ba), ","))

}

func (d deck) shuffle() {

	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
