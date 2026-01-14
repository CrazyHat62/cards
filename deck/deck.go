package deck

import (
	"math/rand/v2"
)

// deafult is "hearts", "diamonds", "clubs", "spades"
var Suits []string = []string{"hearts", "diamonds", "clubs", "spades"}

// default is "A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"
var Values = []string{
	"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K",
}

type Card struct {
	Suit string
	//display string
	Value int
}

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
		for v := range Values {
			var c Card = Card{Suit: s, Value: v}
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

// Color attached to card uses suit to determine if red or black ~ consider generalizing it
func (c Card) Color() string {
	if c.Suit == "hearts" || c.Suit == "diamonds" {
		return "red"
	} else {
		return "black"
	}
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
	println("Pushed", c.Color(), c.Suit, c.Value+1)
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
	//println("Popped", x.Color(), x.suite, x.value+1)
	return x, cards
}

// c, cards := PopLast(cards[:])
func PopLast(cards []Card) (Card, []Card) {
	x := cards[len(cards)-1]
	cards = cards[:len(cards)-1]
	//println("returning", x.Color(), x.suite, x.value+1)
	return x, cards
}

// func printDeck(cards []card) {
// 	println("Deck")
// 	for _, c := range cards {
// 		println(c.Color(), c.suit, c.value+1)
// 	}
// 	print(len(cards))
// }
