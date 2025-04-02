package main

import (
	"math/rand"
)

type Suit struct {
	Name      string
	ShortName string
	Color     string
}

type Rank struct {
	Name      string
	ShortName string
	Value     int
}

type Card struct {
	Suit Suit
	Rank Rank
}

type Deck struct {
	Cards []Card
}

var (
	Hearts   = Suit{Name: "Hearts", ShortName: "\u2665", Color: "Red"}
	Diamonds = Suit{Name: "Diamonds", ShortName: "\u2666", Color: "Red"}
	Clubs    = Suit{Name: "Clubs", ShortName: "\u2663", Color: "Black"}
	Spades   = Suit{Name: "Spades", ShortName: "\u2660", Color: "Black"}
)

var (
	Ace   = Rank{Name: "Ace", ShortName: "A", Value: 1}
	Two   = Rank{Name: "Two", ShortName: "2", Value: 2}
	Three = Rank{Name: "Three", ShortName: "3", Value: 3}
	Four  = Rank{Name: "Four", ShortName: "4", Value: 4}
	Five  = Rank{Name: "Five", ShortName: "5", Value: 5}
	Six   = Rank{Name: "Six", ShortName: "6", Value: 6}
	Seven = Rank{Name: "Seven", ShortName: "7", Value: 7}
	Eight = Rank{Name: "Eight", ShortName: "8", Value: 8}
	Nine  = Rank{Name: "Nine", ShortName: "9", Value: 9}
	Ten   = Rank{Name: "Ten", ShortName: "10", Value: 10}
	Jack  = Rank{Name: "Jack", ShortName: "J", Value: 11}
	Queen = Rank{Name: "Queen", ShortName: "Q", Value: 12}
	King  = Rank{Name: "King", ShortName: "K", Value: 13}
)

func StandardDeck() *Deck {
	deck := &Deck{}

	suits := []Suit{Hearts, Diamonds, Clubs, Spades}
	ranks := []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

	for _, suit := range suits {
		for _, rank := range ranks {
			card := Card{Suit: suit, Rank: rank}
			deck.Cards = append(deck.Cards, card)
		}
	}
	return deck
}

func (d *Deck) Shuffle() {
	for i := len(d.Cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}
