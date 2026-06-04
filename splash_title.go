package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var frames int

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
