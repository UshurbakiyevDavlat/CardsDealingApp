package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")

	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		randInd := rand.Intn(len(d) - 1)

		d[i], d[randInd] = d[randInd], d[i]
	}
}
