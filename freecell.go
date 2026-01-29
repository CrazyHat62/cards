package main

import (
	//"errors"

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

	return err
}

func freeCellScreen(rec rl.Rectangle) {
	var TxSprites rl.Texture2D
	rl.DrawRectangleRec(rec, rl.DarkGreen)

	txt := "press esc to quit"
	txtlen := rl.MeasureText(txt, 30)
	rl.DrawText(txt, screenW/2-txtlen/2-2, screenH/2+2, 30, rl.Black)
	rl.DrawText(txt, screenW/2-txtlen/2, screenH/2, 30, rl.White)

	for i , c := range cards {
		TxSprites = getSpriteSheet(c)
		FrameRecDest.X, FrameRecDest.Y = GetSpritePosition(c)
		Position.X = float32(spriteSizeW * (i % 8)) 
		Position.Y = float32(spriteSizeH * int(i / 8))
		rl.DrawTexturePro(TxSprites, FrameRec, FrameRecDest, Position, 0, rl.RayWhite)
	}


}

func getSpriteSheet(card dk.Card) rl.Texture2D {

	switch {
	case card.Suit == string('C'):
		return ClubsSprites
	case card.Suit == string('S'):
		return SpadesSprites
	case card.Suit == string('D'):
		return DiamondsSprites
	case card.Suit == string('H'):
		return HeartsSprites
	default:
		return PlaceHolderSprites
	}

}

func GetSpritePosition(card dk.Card) (float32, float32) {

	ix := 0
	iy := 0
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
	case card.Rank == string(' '): // blank placeholder
		ix = 0
		iy = 0
	case card.Rank == string('0'): // ace placeholder
		ix = 1
		iy = 0
	}

	return float32(spriteSizeW * ix), float32(spriteSizeH * iy)

}
