package main

import (
	"math/rand/v2"
)

// deafult is "hearts", "diamonds", "clubs", "spades"
var Suits []string = []string{"hearts", "diamonds", "clubs", "spades"}

// default is "A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"
var Ranks = []string{
	"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K",
}

var Values = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

type Card struct {
	Suit  string
	Rank  string
	value int
}

// Color attached to card uses suit to determine if red or black ~ consider generalizing it
func (c Card) Color() string {
	if c.Suit == "hearts" || c.Suit == "diamonds" {
		return "red"
	} else {
		return "black"
	}
}
func (c Card) CardValue() int {
	return c.value
}

//func (c Card) Card
// func main() {
// 	cards := FreshDeck()
// 	printDeck(cards)
// 	c, cards := PopFirst(cards[:])
// 	cards = PushLast(cards[:], c)
// 	printDeck(cards)
// }

// FreshDeck is an unshuffled new deck of cards
func FreshDeck() []Card {
	var cards []Card
	for _, s := range Suits {

		for i, v := range Ranks {
			var c Card = Card{Suit: s, Rank: v, value: Values[i]}
			cards = append(cards, c)
		}
	}

	return cards
}

// ShuffledDeck is a convienience function to grab a shuffled fresh deck
func ShuffledDeck() []Card {
	cards := FreshDeck()
	Shuffle(cards[:])
	return cards
}

// Shuffle changes the deck you passed - pass by slice
// e.g. Shuffle(cards[:])
func Shuffle(cards []Card) {
	for i := len(cards) - 1; i > 0; i-- {
		j := rand.IntN(len(cards))
		cards[i], cards[j] = cards[j], cards[i]
	}
}

// PushLast will place the card at the bottom of the cards
// e.g. cards = PushLast(cards[:], c)
func PushLast(cards []Card, c Card) []Card {
	cards = append(cards, c)
	println("Pushed", c.Color(), c.Suit, c.CardValue())
	return cards
}

// PushFirst will place the card at the Top of the cards
// e.g. cards = PushFirst(cards[:], c)
func PushFirst(cards []Card, c Card) []Card {
	cards = append([]Card{c}, cards...)
	return cards
}

// e.g. c, cards := PopFirst(cards[:])
func PopFirst(cards []Card) (Card, []Card) {
	x := cards[0]
	cards = cards[1:]
	//println("Popped", x.Color(), x.suite, x.CardValue())
	return x, cards
}

// c, cards := PopLast(cards[:])
func PopLast(cards []Card) (Card, []Card) {
	x := cards[len(cards)-1]
	cards = cards[:len(cards)-1]
	//println("returning", x.Color(), x.suite, x.CardValue())
	return x, cards
}

// func printDeck(cards []card) {
// 	println("Deck")
// 	for _, c := range cards {
// 		println(c.Color(), c.suit, c.CardValue())
// 	}
// 	print(len(cards))
// }
