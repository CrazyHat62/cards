package main

import (
	//"errors"
	"fmt"
	"math/rand"

	dk "example.com/deck"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var cards []dk.Card
var card dk.Card

// var hand []Card
var CurrentGameSeed int
var deckPlayed bool

func freeCellGame() error {

	if !deckPlayed {
		deckPlayed = true
		seed := 1 + rand.Intn(32000)
		cards = dk.Deal(seed)
	}

	if rl.IsKeyPressed(rl.KeyEnter) {
		card, cards, err = dk.PopFirst(cards[:])
		// 	seed := 1 + rand.Intn(32000)
		// 	cards = deal(seed)
		// 	//Shuffle(cards[:])
	}
	return err
}

func freeCellScreen(rec rl.Rectangle) {
	var TxSprites rl.Texture2D
	rl.DrawRectangleRec(rec, rl.DarkGreen)

	txt := "The card you drew is " + fmt.Sprint(card.Rank) + " of " + card.Suit
	txtlen := rl.MeasureText(txt, 50)
	rl.DrawText(txt, screenW/2-txtlen/2-2, screenH/2-50+2, 50, rl.Black)
	rl.DrawText(txt, screenW/2-txtlen/2, screenH/2-50, 50, rl.White)
	txt = "press enter to move to next screen"
	txtlen = rl.MeasureText(txt, 30)
	rl.DrawText(txt, screenW/2-txtlen/2-2, screenH/2+2, 30, rl.Black)
	rl.DrawText(txt, screenW/2-txtlen/2, screenH/2, 30, rl.White)

	//rl.DrawTextureRec(TxSprites, FrameRec, Position, rl.RayWhite)
	switch {

	case card.Suit == string('C'):
		TxSprites = ClubsSprites
	case card.Suit == string('S'):
		TxSprites = SpadesSprites
	case card.Suit == string('D'):
		TxSprites = DiamondsSprites
	case card.Suit == string('H'):
		TxSprites = HeartsSprites
	}
	ix := int32(0)
	iy := int32(0)
	switch { //could use dictionary to map
	case card.Rank == string('2'):
		ix = 0
		iy = 0
	case card.Rank == string('3'):
		ix = 1
		iy = 0
	case card.Rank == string('4'):
		ix = 2
		iy = 0
	case card.Rank == string('5'):
		ix = 3
		iy = 0
	case card.Rank == string('6'):
		ix = 4
		iy = 0
	case card.Rank == string('7'):
		ix = 5
		iy = 0
	case card.Rank == string('8'):
		ix = 6
		iy = 0
	case card.Rank == string('9'):
		ix = 7
		iy = 0
	case card.Rank == string('T'):
		ix = 0
		iy = 1
	case card.Rank == string('A'):
		ix = 1
		iy = 1
	case card.Rank == string('J'):
		ix = 5
		iy = 1
	case card.Rank == string('K'):
		ix = 6
		iy = 1
	case card.Rank == string('Q'):
		ix = 7
		iy = 1
	}
	FrameRec.X = float32(spriteSizeW * ix) // + 2*ix) padding??
	FrameRec.Y = float32(spriteSizeH * iy)

	rl.DrawTexturePro(TxSprites, FrameRec, FrameRecDest, Position, 0, rl.RayWhite)

}
