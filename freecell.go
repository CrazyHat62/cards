package main

import (
	"fmt"
	"math/rand/v2"

	dk "example.com/deck"
	rl "github.com/gen2brain/raylib-go/raylib"
)

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

	return nil
}

func updateFreeCell() {
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
}

// **** Freecell
func freeCellScreen(rec rl.Rectangle) {
	var TxSprites rl.Texture2D
	rl.DrawRectangleRec(rec, rl.DarkGreen)

	for i, card := range cards {
		TxSprites = dk.GetSuitSprite(card, TxSprites)
		origin := rl.Vector2(floatVect{X: 0.0, Y: 0.0})
		rotation := float32(0.0)
		if !card.IsSelected {
			rl.DrawTexturePro(TxSprites, cards[i].Source, cards[i].Dest, origin, rotation, rl.RayWhite)
		}
	}
	// Draw selected card(s) last so they appear on top (use current Dest which may be moved while dragging)
	for _, card := range cards {
		if card.IsSelected {
			TxSprites = dk.GetSuitSprite(card, TxSprites)
			origin := rl.Vector2(floatVect{X: 0.0, Y: 0.0})
			rotation := float32(0.0)
			rl.DrawTexturePro(TxSprites, card.Source, card.Dest, origin, rotation, rl.RayWhite)
		}
	}
}
