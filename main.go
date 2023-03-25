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
	card := suffle.GetCard("King of Clubs")
	cards := []types.Card{
		suffle.GetCard("Ace of Hearts"),
		suffle.GetCard("Jack of Spades"),
		suffle.GetCard("Queen of Diamonds"),
	}

	fmt.Println(card)
	fmt.Println(cards)
	print.Scan("well done")
}