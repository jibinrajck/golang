package main

func main() {
	//type deck []string
	cards := newDeck()
	hand, remainingcard := deal(cards, 5)
	hand.printcards()
	remainingcard.printcards()
	println(cards.toString())
}

func newCard() string {
	return "Heart"
}
