package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	// cards := deck{}
	var cards deck

	cardSuits := []string{"Space", " Diadmonds", " Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, card := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+card)
		}
	}
	return cards
}

// It is an receiver function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(cardstack deck, handSize int) (deck, deck) {
	return cardstack[:handSize], cardstack[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// save to a file which is present on a hardware
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// extact from a file
func newDeckFromFile(filename string) deck {
	byte_ans, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error occured while reading a file:", err)
		os.Exit(1)
	}
	finalans := strings.Split(string(byte_ans), ",")
	return deck(finalans)
}

// now to shuffle the cards available in our deck we need to go through some specific way
// since in GO we dont have inbuilt library to shuffle the pattern.
// one way :
//for each index , card in cards:
//		Generate a random number between 0 to len(cards)-1
//		swap the current card and card at cards[random number]

func (d deck) shuffle() {

	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
