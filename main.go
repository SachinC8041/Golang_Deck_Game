package main

import (
	"fmt"
)

func main() {
	fmt.Println("Card Deck Game")
	cards := newDeckFromFile("deckgame.txt")
	// firstHand, secondHand := deal(cards, 5)
	// firstHand.print()
	// fmt.Println("-------------")
	// secondHand.print()
	// fmt.Println("--------------")
	// fmt.Println(cards.toString())
	// fmt.Println("__------------------------")
	// cards.saveToFile("deckgame.txt")
	cards.shuffle()
	cards.print()
}
