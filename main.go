package main

func main() {
	//type deck []string
	cards := newDeck()
	cards.printcards()
}

func newCard() string {
	return "Heart"
}
