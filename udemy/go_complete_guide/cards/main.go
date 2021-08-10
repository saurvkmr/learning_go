package main

import "fmt"

var name string

func main() {
	cardArray := [3]string{"Ace of Spades", newCard()}
	card := []string{"Ace of Spades", newCard()}

	fmt.Println(cardArray)
	fmt.Println(card)
}

func newCard() string {
	return "Ace of Dimond"
}
