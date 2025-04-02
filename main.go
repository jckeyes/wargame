package main

import (
	"fmt"
)

func main() {
	d := StandardDeck()
	d.Shuffle()

	for _, card := range d.Cards {
		// println("${card.Rank.ShortName} of ${card.Suit.ShortName}")
		fmt.Printf("%s%s\n", card.Rank.ShortName, card.Suit.ShortName)
	}
}
