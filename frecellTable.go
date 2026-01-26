package main

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
