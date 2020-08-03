package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Create type deck : slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, value := range cardValues {
		for _, suit := range cardSuits {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

// (d deck) is a receiver. any var of type deck can access print method
// d is a copy (or ref ?) we are working with
// convention d is because we take first 1 or 2 or 3 letters abbr. of the type
func (d deck) print() {
	// redeclaring i card for each loop
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) deal2(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	data := []byte(d.toString())
	return ioutil.WriteFile(filename, data, 0666)
}
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	// fmt.Println(err)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// OPtion #2 - Log the err and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}
