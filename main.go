package main

import (
	"errors"
	"fmt"
	"math/rand"

	//deck "example.com/deck"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const screenW = int32(1280)
const screenH = int32(720)

const (
	LOGO = iota
	TITLE
	GAMEPLAY
	ENDING
)

var splashCountDown int = 240
var cards []Card
var card Card

// var hand []Card
var CurrentGameSeed int
var deckPlayed bool
var err error

var frames int

const spriteSizeW = 125 //500
const spriteSizeH = 181 //726

var ClubsSprites rl.Texture2D
var SpadesSprites rl.Texture2D
var DiamondsSprites rl.Texture2D
var HeartsSprites rl.Texture2D
var FrameRec rl.Rectangle
var FrameRecDest rl.Rectangle
var Position rl.Vector2

// type gameScreen int
type gameScreen func(rl.Rectangle)

type game func() error

func main() {

	rl.InitWindow(screenW, screenH, "raylib [core] example - basic screen manager")

	ClubsSprites = rl.LoadTexture("images/clubsS.png")
	defer rl.UnloadTexture(ClubsSprites)
	SpadesSprites = rl.LoadTexture("images/spadesS.png")
	defer rl.UnloadTexture(SpadesSprites)
	DiamondsSprites = rl.LoadTexture("images/diamondsS.png")
	defer rl.UnloadTexture(DiamondsSprites)
	HeartsSprites = rl.LoadTexture("images/heartsS.png")
	defer rl.UnloadTexture(HeartsSprites)

	var currentScreen gameScreen
	var currentGame game
	currentScreen = splashScreen
	currentGame = noGame
	currentScreenENUM := LOGO
	card = noCard()

	FrameRec = rl.NewRectangle(0, 0, spriteSizeW, spriteSizeH)
	FrameRecDest = rl.NewRectangle(0, 0, float32(spriteSizeW/2), float32(spriteSizeH/2))
	Position = rl.NewVector2(0.0, 0.0)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		switch currentScreenENUM {
		case LOGO:
			if frames > splashCountDown {
				currentScreenENUM = TITLE
				currentScreen = titleScreen
				currentGame = noGame
				card = noCard()
			}
		case TITLE:
			if rl.IsKeyPressed(rl.KeyEnter) {
				currentScreenENUM = GAMEPLAY
				currentScreen = freeCellScreen
				currentGame = freeCellGame
			}
		case GAMEPLAY:
			if rl.IsKeyPressed(rl.KeyEscape) || (deckPlayed && card.Suit == "") {
				deckPlayed = false
				currentScreenENUM = ENDING
				currentScreen = endScreen
				currentGame = noGame
				card = noCard()
			}
		case ENDING:
			if rl.IsKeyPressed(rl.KeyEnter) {

				//currentScreenENUM = LOGO
				//currentScreen = splashScreen
				//currentGame = noGame
				//frames = 0 //reset count down for logo
				currentScreenENUM = TITLE
				currentScreen = titleScreen
				currentGame = noGame
				card = noCard()
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.White)
		rec := rl.NewRectangle(0, 0, float32(screenW), float32(screenH))

		err = currentGame()
		currentScreen(rec)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func noGame() error {
	return errors.New("no game")
}

func freeCellGame() error {

	if !deckPlayed {
		deckPlayed = true
		seed := 1 + rand.Intn(32000)
		cards = deal(seed)
	}

	if rl.IsKeyPressed(rl.KeyEnter) {
		card, cards, err = PopFirst(cards[:])
		// 	seed := 1 + rand.Intn(32000)
		// 	cards = deal(seed)
		// 	//Shuffle(cards[:])
	}
	return err
}

// func tonysGameUI(rec rl.Rectangle) {
// 	_ = deck.FreshDeck()

// 	rl.DrawRectangleRec(rec, rl.DarkPurple)
// }

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
	}
	FrameRec.X = float32(spriteSizeW * ix) // + 2*ix) padding??
	FrameRec.Y = float32(spriteSizeH * iy)

	rl.DrawTexturePro(TxSprites, FrameRec, FrameRecDest, Position, 0, rl.RayWhite)

}

func splashScreen(rec rl.Rectangle) {
	frames++
	txt := "YOUR LOGO GOES HERE"
	txtlen := rl.MeasureText(txt, 50)
	rl.DrawText(txt, screenW/2-txtlen/2-3, screenH/2-50+3, 50, rl.Magenta)
	rl.DrawText(txt, screenW/2-txtlen/2-1, screenH/2-50+1, 50, rl.Black)
	rl.DrawText(txt, screenW/2-txtlen/2, screenH/2-50, 50, rl.White)
	txt = "this message disappears in " + fmt.Sprint(splashCountDown-frames) + " frames"
	txtlen = rl.MeasureText(txt, 30)
	rl.DrawText(txt, screenW/2-txtlen/2-3, screenH/2+3, 30, rl.Magenta)
	rl.DrawText(txt, screenW/2-txtlen/2-1, screenH/2+1, 30, rl.Black)
	rl.DrawText(txt, screenW/2-txtlen/2, screenH/2, 30, rl.White)
}

func titleScreen(rec rl.Rectangle) {
	rl.DrawRectangleRec(rec, rl.DarkPurple)
	txt := "AN AMAZING TITLE GOES HERE"
	txtlen := rl.MeasureText(txt, 50)
	rl.DrawText(txt, screenW/2-txtlen/2-2, screenH/2-50+2, 50, rl.Black)
	rl.DrawText(txt, screenW/2-txtlen/2, screenH/2-50, 50, rl.White)
	txt = "press enter to move to next screen"
	txtlen = rl.MeasureText(txt, 30)
	rl.DrawText(txt, screenW/2-txtlen/2-2, screenH/2+2, 30, rl.Black)
	rl.DrawText(txt, screenW/2-txtlen/2, screenH/2, 30, rl.White)
}

func endScreen(rec rl.Rectangle) {
	rl.DrawRectangleRec(rec, rl.DarkBlue)
	txt := "A DRAMATIC ENDING GOES HERE"
	txtlen := rl.MeasureText(txt, 50)
	rl.DrawText(txt, screenW/2-txtlen/2-2, screenH/2-50+2, 50, rl.Black)
	rl.DrawText(txt, screenW/2-txtlen/2, screenH/2-50, 50, rl.White)
	txt = "press enter to move to next screen"
	txtlen = rl.MeasureText(txt, 30)
	rl.DrawText(txt, screenW/2-txtlen/2-2, screenH/2+2, 30, rl.Black)
	rl.DrawText(txt, screenW/2-txtlen/2, screenH/2, 30, rl.White)
}
