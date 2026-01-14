module example.com/cards

go 1.25.5

//replace example.com/deck => ./deck/

//example.com/deck v0.0.0
require github.com/gen2brain/raylib-go/raylib v0.55.1

require (
	github.com/ebitengine/purego v0.9.1 // indirect
	golang.org/x/exp v0.0.0-20260112195511-716be5621a96 // indirect
	golang.org/x/sys v0.40.0 // indirect
)
