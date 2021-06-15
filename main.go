package main

func main() {
	cards := newDeck()
	cards.print()
	cards.saveToFile("my_file")

	cards2 := newDeckFromFile("my_file")
	cards2.print()

	cards2.shuffle2()
	cards2.print()
}
