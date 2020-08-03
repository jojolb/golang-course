package main

import "fmt"

var outVar int

func main() {
	// var card string = "Ace of spades"
	// Same as above, go infers type from right side value when using :=
	card := "Ace of spades"

	card = "Five of diamonds"

	var noInit string

	noInit = "Jojo"
	outVar = 3
	fmt.Println(card)
	fmt.Println(noInit)

	cards := []string{newCard(), "Ace of Spades"}

	cards = append(cards, "six of Spades")

	fmt.Println(cards)

	// redeclaring i card for each loop
	for i, card := range cards {
		fmt.Println(i, card)
	}

	// type conversion
	greeting := "Hi There!"
	fmt.Println([]byte(greeting))
}

func newCard() string {
	return "Five of Diamonds"
}
