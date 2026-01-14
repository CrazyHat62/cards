package main

import (
	"testing"
)

func TestFreshDeck(t *testing.T) {

	cards := FreshDeck()
	suits := Suits
	ranks := Ranks
	suit := 0

	for i, c := range cards {
		v := i % len(ranks)

		if c.Suit != suits[suit] || c.value != v + 1 {

			t.Errorf("got %q %d want %q %d", c.Suit, c.value, suits[suit], v+1)
		}

		//we up the suit after being through all the values
		if v == (len(ranks) - 1) {
			suit++
		}
		//println(c.Color(), c.suite, c.value)
	}
	print(len(cards))

}

func TestShuffle(t *testing.T) {
	cardsW := FreshDeck()
	cards := FreshDeck()
	Shuffle(cards[:])
	//length must be the same
	if len(cards) != len(cardsW) {
		t.Errorf("got %d want %d", len(cards), len(cardsW))
		return
	}
	//shuffle order must be different to sorted
	count := 0
	atleast := 1
	for i, c := range cardsW {
		cg := cards[i]
		if c.Suit != cg.Suit || c.CardValue() != cg.CardValue() {
			count++
		}
	}
	if count < atleast {
		t.Errorf("got %d want > %d", count, atleast)
	}
	//all the cards in wanted must exist in got
OuterLoop:
	for _, c := range cardsW {
		for _, cg := range cards {
			if c.Suit == cg.Suit && c.value == cg.value {
				continue OuterLoop
			}
		}
		t.Errorf("not found: wanted %q %q", c.value, c.Suit)
	}
	//finally the next shuffle cant be the same as the first
	Shuffle(cardsW[:])
	count = 0
	atMost := 5
	for i, c := range cardsW {
		cg := cards[i]
		if c.Suit == cg.Suit && c.value == cg.value {
			count++
		}
	}
	if count > atMost {
		t.Errorf("got %d want > %d", count, atMost)
	}

}

func TestColor(t *testing.T) {
	var c Card = Card{Suit: "hearts", value: 0}
	if c.Color() != "red" {
		t.Errorf("got %q want %q", c.Color(), "red")
	}
	c = Card{Suit: "spades", value: 13}
	if c.Color() != "black" {
		t.Errorf("got %q want %q", c.Color(), "black")
	}

}
