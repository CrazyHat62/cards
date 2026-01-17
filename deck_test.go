package main

import (
	"math/rand"
	"testing"
)

func TestDeal(t *testing.T) {

	got := deal(2570)
	want := []string{"2S", "KD", "JD", "8H", "8S", "QD", "4C", "3S",
		"8D", "AH", "3C", "5S", "2D", "7C", "7D", "QC",
		"QS", "6C", "6D", "5H", "KH", "KS", "7S", "9H",
		"4D", "2H", "3D", "9C", "4S", "QH", "7H", "6S",
		"AC", "3H", "KC", "TD", "2C", "4H", "JC", "TH",
		"9S", "5D", "JH", "TC", "6H", "9D", "JS", "AS",
		"5C", "8C", "TS", "AD"}

	for i, c := range got {

		g := c.Rank + c.Suit
		if g != want[i] {
			t.Errorf("got %q want %q", g, want[i])
		}

	}
	print(len(cards))

}

func TestShuffle(t *testing.T) {

	gotSeed := 1 + rand.Intn(32000)
	wantSeed := 1 + rand.Intn(32000)

	got := deal(gotSeed)
	want := deal(wantSeed)
	//length must be the same
	if len(got) != len(want) {
		t.Errorf("got %d want %d", len(got), len(want))
		return
	}
OuterLoop:
	for _, w := range want {
		for _, g := range got {
			if w.Suit == g.Suit && w.Rank == g.Rank {
				continue OuterLoop
			}
		}
		t.Errorf("not found: wanted %q %q", w.Rank, w.Suit)
	}

	count := 0
	atMost := 5
	for i, c := range want {
		cg := got[i]
		if c.Suit == cg.Suit && c.index == cg.index {
			count++
		}
	}
	if count > atMost {
		t.Errorf("got %d want > %d", count, atMost)
	}

}

func TestColor(t *testing.T) {
	var c Card = Card{Suit: "hearts", index: 0}
	if c.Color() != "red" {
		t.Errorf("got %q want %q", c.Color(), "red")
	}
	c = Card{Suit: "spades", index: 13}
	if c.Color() != "black" {
		t.Errorf("got %q want %q", c.Color(), "black")
	}

}
