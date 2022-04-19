package main

import (
	"fmt"
	"math/rand"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "club", "Heart", "Diamond", "Ruby"}
	cardValue := []string{"Ace", "two", "three", "Four", "five", "six", "seven", "Eight", "nine ", "ten", "king", "queen", "jack"}

	for _, suit := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	if handSize < 0 {
		fmt.Println("this is not going to work")

	}
	return d[:handSize], d[handSize:]

}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newposition := r.Intn(len(d) - 1)

		d[i], d[newposition] = d[newposition], d[i]

	}
}
