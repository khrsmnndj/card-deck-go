package main

import (
	"fmt"
	"github.com/khrsmnndj/card-deck-go/src/print"
	"github.com/khrsmnndj/card-deck-go/src/suffle"
	"github.com/khrsmnndj/card-deck-go/src/suffle/types"
)

func main(){
	fmt.Println("Card Deck Project")
	fmt.Println("The Card")
	var cards types.Cards
	cards = append(cards,
		suffle.GetCard("Ace of Hearts"),
		suffle.GetCard("Jack of Spades"),
		suffle.GetCard("Queen of Diamonds"),
	)
	cards.Append("Spades of ", 7)
	cards.Append("Clubs of ", 7)
	fmt.Println(cards)
	print.Scan("well done")
}