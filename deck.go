package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string // << Create a custom type called deck

func newDeck() deck {
	cards := deck{"Ace of Spades", "Two of Hearts", "Three of Diamonds"}
	return cards
}

func (d deck) print() { // << Define receiver function for deck type
	for i, card := range d {
		fmt.Println(i, " ", card)
	}
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
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

// func (d deck) shuffle() {
// 	for i := range d {
// 		newPosition := rand.Intn(len(d) - 1)

// 		d[i], d[newPosition] = d[newPosition], d[i]
// 	}
// }

func (d deck) shuffle2() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
