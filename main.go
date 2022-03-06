package main

import "fmt"

func main() {
	//type deck []string
	cards := newDeck()
	hand, remainingcard := deal(cards, 5)
	hand.printcards()
	remainingcard.printcards()
	println(cards.toString())
	cards.saveToFile("FirstDeck")
	newDeckReadFromFile := newDeckFromFile("FirstDeck")
	fmt.Println("From File   : ", newDeckReadFromFile)

	cards.shuffle()

	fmt.Println("Shuffled Cards")
	cards.printcards()
}

func newCard() string {
	return "Heart"
}
