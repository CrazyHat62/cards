module example.com/cards

go 1.25.5

replace example.com/deck => ./deck/

require (
	example.com/deck v0.0.0
	github.com/gen2brain/raylib-go/raylib v0.60.0
)

require (
	github.com/ebitengine/purego v0.10.0 // indirect
	github.com/jupiterrider/ffi v0.7.0 // indirect
	golang.org/x/exp v0.0.0-20260508232706-74f9aab9d74a // indirect
)
