package main

import (
	"fmt"
)

func main() {
	cards := newDeck()

	// cards.print()
	hand := deck{}
	remainingCards := deck{}

	// Both work but course is using 1st
	// hand, remainingCards := deal(cards, 5)
	hand, remainingCards = cards.deal2(5)
	hand.print()
	remainingCards.print()
	// fmt.Println(cards.toString())

	fmt.Println("Save : ", cards.saveToFile("saved"))
	cards = newDeckFromFile("saved")
	cards.print()
}
