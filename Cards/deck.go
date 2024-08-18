package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	//init list for suits
	cardOfSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}

	//init list for values
	cardOfValues := []string{"One", "Two", "Three", "Four"}

	//iterate them across each other and create combinations
	for _, suit := range cardOfSuits {
		for _, value := range cardOfValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	//return deck
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
