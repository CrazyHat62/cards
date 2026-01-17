package main

import (
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
var hand []Card
var CurrentGameSeed int
var frames int

// type gameScreen int
type gameScreen func(rl.Rectangle)
type game func()

func main() {

	rl.InitWindow(screenW, screenH, "raylib [core] example - basic screen manager")

	var currentScreen gameScreen
	var currentGame game
	currentScreen = splashScreen
	currentGame = noGame
	currentScreenENUM := LOGO

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		switch currentScreenENUM {
		case LOGO:
			if frames > splashCountDown {
				currentScreenENUM = TITLE
				currentScreen = titleScreen
			}
		case TITLE:
			if rl.IsKeyPressed(rl.KeyEnter) {
				currentScreenENUM = GAMEPLAY
				currentScreen = freeCellScreen
				currentGame = freeCellGame
			}
		case GAMEPLAY:
			if rl.IsKeyPressed(rl.KeyEnter) {
				currentScreenENUM = ENDING
				currentScreen = endScreen
				currentGame = noGame
			}
		case ENDING:
			if rl.IsKeyPressed(rl.KeyEnter) {
				currentScreenENUM = LOGO
				currentScreen = splashScreen
				//				frames = 0 //reset count down for logo if looping
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rec := rl.NewRectangle(0, 0, float32(screenW), float32(screenH))

		currentGame()
		currentScreen(rec)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func noGame() {

}

func freeCellGame() {
	seed := 1 + rand.Intn(32000)
	cards = deal(seed)
	hand = nil
	Shuffle(cards[:])
	card, cards = PopFirst(cards[:])
	hand = append(hand, card)

}

// func tonysGameUI(rec rl.Rectangle) {
// 	_ = deck.FreshDeck()

// 	rl.DrawRectangleRec(rec, rl.DarkPurple)
// }

func freeCellScreen(rec rl.Rectangle) {
	rl.DrawRectangleRec(rec, rl.DarkPurple)

	txt := "The card you drew is " + fmt.Sprint(hand[0].Rank) + " of " + hand[0].Suit
	txtlen := rl.MeasureText(txt, 50)
	rl.DrawText(txt, screenW/2-txtlen/2-2, screenH/2-50+2, 50, rl.Black)
	rl.DrawText(txt, screenW/2-txtlen/2, screenH/2-50, 50, rl.White)
	txt = "press enter to move to next screen"
	txtlen = rl.MeasureText(txt, 30)
	rl.DrawText(txt, screenW/2-txtlen/2-2, screenH/2+2, 30, rl.Black)
	rl.DrawText(txt, screenW/2-txtlen/2, screenH/2, 30, rl.White)

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
	rl.DrawRectangleRec(rec, rl.DarkGreen)
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
