package main

import "fmt"

func main() {
	cards := newDeck()

	cards.shuffle()
	cards.print()
	fmt.Println()

	hand, remaingcrads := deal(cards, 5)
	hand.print()
	fmt.Println()
	remaingcrads.print()

}
