// This app is a freecell solitare game

package main

import (
	"errors"
	"fmt"
	"os"

	"math/rand/v2"

	dk "example.com/deck"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var screenW int32 = 1280
var screenH int32 = 720

const (
	LOGO = iota
	TITLE
	GAMEPLAY
	ENDING
)

var splashCountDown int = 240

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

var cards []dk.Card
var card dk.Card

// var hand []Card
var CurrentGameSeed int
var deckPlayed bool
var Suits []string = []string{"hearts", "diamonds", "clubs", "spades"}
var Ranks []string = []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
var Values = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

// Mouse dragging variables
var dragIndex int = -1
var dragOffset rl.Vector2

// Start of functions
//
//
//

// type gameScreen
type gameScreen func(rl.Rectangle)

type game func() error

func main() {

	//rl.InitWindow(0, 0, "raylib [core] example - basic screen manager") //fullscreen
	rl.InitWindow(screenW, screenH, "raylib [core] example - basic screen manager")
	defer rl.CloseWindow()
	screenW = int32(rl.GetScreenWidth())
	screenH = int32(rl.GetScreenHeight())

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
	card = dk.NoCard()
	//card = NoCard()

	FrameRec = rl.NewRectangle(0, 0, spriteSizeW, spriteSizeH)
	FrameRecDest = rl.NewRectangle(0, 0, float32(spriteSizeW/2), float32(spriteSizeH/2))
	Position = rl.NewVector2(0.0, 0.0)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		// --- Update Logic ---
		switch currentScreenENUM {
		case GAMEPLAY:
			// Mouse press: pick topmost card under cursor and begin dragging
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				mousePos := rl.GetMousePosition()

				for i := len(cards) - 1; i >= 0; i-- {
					if rl.CheckCollisionPointRec(mousePos, cards[i].Dest) {
						fmt.Printf("\nFound card at %v\n", cards[i].Dest)
						dragIndex = i
						cards[i].IsSelected = true
						dragOffset = rl.NewVector2(mousePos.X-cards[i].Dest.X, mousePos.Y-cards[i].Dest.Y)
						break
					}
				}
			}

			// While holding mouse, move the dragged card with the cursor
			if dragIndex != -1 && rl.IsMouseButtonDown(rl.MouseLeftButton) {
				mousePos := rl.GetMousePosition()
				cards[dragIndex].Dest.X = mousePos.X - dragOffset.X
				cards[dragIndex].Dest.Y = mousePos.Y - dragOffset.Y
			}

			// On release, drop the card
			if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
				if dragIndex != -1 {
					cards[dragIndex].IsSelected = false
					dragIndex = -1
				}
			}

			if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyEnter) { //|| (deckPlayed && card.Suit == "") {
				deckPlayed = false
				currentScreenENUM = ENDING
				currentScreen = endScreen
				currentGame = noGame
				card = dk.NoCard()

			}
		case LOGO:
			if frames > splashCountDown {
				currentScreenENUM = TITLE
				currentScreen = titleScreen
				currentGame = noGame
				card = dk.NoCard()
				//card = NoCard()
			}
		case TITLE:
			if rl.IsKeyPressed(rl.KeyEnter) {
				currentScreenENUM = GAMEPLAY
				// Set up free cell here
				currentScreen = freeCellScreen
				currentGame = freeCellGame
			}
		case ENDING:
			if rl.IsKeyPressed(rl.KeyEnter) {
				os.Exit(0)
			}
		}

		// --- Drawing Logic ---
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rec := rl.NewRectangle(0, 0, float32(screenW), float32(screenH))

		// get deck and card
		err = currentGame()
		// draw card(s)
		currentScreen(rec)

		rl.EndDrawing()
	}
}

func noGame() error {
	return errors.New("no game")
}

// func tonysGameUI(rec rl.Rectangle) {
// 	_ = deck.FreshDeck()

// 	rl.DrawRectangleRec(rec, rl.DarkPurple)
// }

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

// **** Freecell
func freeCellScreen(rec rl.Rectangle) {
	var TxSprites rl.Texture2D
	rl.DrawRectangleRec(rec, rl.DarkGreen)

	for i, card := range cards {
		TxSprites = getSuitSprite(card, TxSprites)
		origin := rl.Vector2(floatVect{X: 0.0, Y: 0.0})
		rotation := float32(0.0)
		if !card.IsSelected {
			rl.DrawTexturePro(TxSprites, cards[i].Source, cards[i].Dest, origin, rotation, rl.RayWhite)
		}
	}
	// Draw selected card(s) last so they appear on top (use current Dest which may be moved while dragging)
	for _, card := range cards {
		if card.IsSelected {
			TxSprites = getSuitSprite(card, TxSprites)
			origin := rl.Vector2(floatVect{X: 0.0, Y: 0.0})
			rotation := float32(0.0)
			rl.DrawTexturePro(TxSprites, card.Source, card.Dest, origin, rotation, rl.RayWhite)
		}
	}
}

//var cards []dk.Card
//var card dk.Card

func freeCellGame() error {

	if !deckPlayed {
		deckPlayed = true
		seed := 1 + rand.IntN(32000)
		cards = dk.Deal(seed)
		for i, card := range cards {
			cards[i].Source = getCardSource(card, FrameRec)
			cards[i].Dest = getInitialGridDest(i, FrameRec)
		}
	}

	//if rl.IsKeyPressed(rl.KeyEnter) {
	//	card, cards, err = dk.PopFirst(cards[:])
	// 	seed := 1 + rand.Intn(32000)
	// 	cards = deal(seed)
	// 	//Shuffle(cards[:])
	//}
	return nil
}

// returns the source rectangle for the card sprite sheet based on the card's rank and suit
func getCardSource(card dk.Card, frameRec rl.Rectangle) rl.Rectangle {
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

	source := frameRec
	source.X = float32(spriteSizeW*ix + 1)
	source.Y = float32(spriteSizeH * iy)

	return source
}

func getInitialGridDest(idx int, frameRec rl.Rectangle) rl.Rectangle {
	dest := frameRec
	xgrd := idx % 8
	ygrd := idx / 8
	dest.X = float32(spriteSizeW*xgrd + 10*(xgrd+1) + 90)
	dest.Y = float32(spriteSizeH + (ygrd)*50 + 40)
	return dest
}

func getSuitSprite(card dk.Card, TxSprites rl.Texture2D) rl.Texture2D {
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
	return TxSprites
}

// *** FreecellTable

type Grid struct {
	NumRows, NumCols, HolderSize int
	gridData                     []int
	//Colors                     []rl.Color
}

// Initilize all values with zeros
func NewGrid() *Grid {
	g := new(Grid)
	//if g.NumRows <= 0 || g.NumCols <= 0 {
	g.NumRows = 20
	g.NumCols = 10
	g.HolderSize = 30

	//}
	if len(g.gridData) == 0 {
		g.gridData = make([]int, g.NumRows*g.NumCols)
	}
	for i := range g.gridData {
		g.gridData[i] = 0
	}
	//	g.Colors = g.getCellColors()
	return g
}

// **** AspectRatio
type floatVect struct{ X, Y float32 }

// func getAspectRatio(a, b int32) floatVect {
// 	var GCD int32

// 	GCD = gcd(screenW, screenH)
// 	w := screenW / GCD
// 	h := screenH / GCD
// 	arw := float32(w) / float32(h) //take note '/' does not work the same as in C
// 	return floatVect{X: arw, Y: 1.0}
// }

// // gcd (Greatest Common Divisor) calculates the GCF of two numbers using the Euclidean algorithm.
// func gcd(a, b int32) int32 {
// 	// Base case: if the second number (b) is 0, the GCD is the first number (a).
// 	if b == 0 {
// 		return a
// 	}
// 	// Recursive step: call gcd with the second number (b) and the remainder of a divided by b.
// 	return gcd(b, a%b)
// }
