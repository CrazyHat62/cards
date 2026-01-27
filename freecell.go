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
	source := FrameRec
	source.X = float32(spriteSizeW*ix + ix)
	source.Y = float32(spriteSizeH*iy + iy)
	dest := FrameRec //FrameRecDest
	dest.X = float32(spriteSizeW*ix+ix) + 10.0
	dest.Y = float32(spriteSizeH*iy+iy) + 10.0

	origin := rl.Vector2(floatVect{X: 0.0, Y: 0.0})
	rotation := float32(0.0)

	rl.DrawTexturePro(TxSprites, source, dest, origin, rotation, rl.RayWhite)

}

/*

+ texture: The Texture2D to be drawn. This must be loaded into GPU memory first using a function like LoadTexture.

+ source: A Rectangle that defines the area of the original texture to use (in texture space, where (0, 0) is the top-left corner of the texture).

    To use the entire texture, the source rectangle would be { 0.0f, 0.0f, (float)texture.width, (float)texture.height }.
    To flip the texture horizontally, use a negative width in the source rectangle.

+ dest: A Rectangle that defines where on the screen the source rectangle will be drawn (in screen space, where (0, 0) is the top-left corner of the window). The width and height of this rectangle determine the final scale of the drawn texture portion.

+ origin: A Vector2 that specifies the point around which rotation and scaling will occur. This point is relative to the top-left corner of the destination rectangle (dest).

    For no rotation/custom origin, use { 0.0f, 0.0f }.
    To rotate around the center of the destination rectangle, use { dest.width / 2, dest.height / 2 }.

+ rotation: A float value for the angle of rotation in degrees (clockwise).

+ tint: A Color to apply as a tint to the texture. Using WHITE displays the texture in its native colors.

*/
