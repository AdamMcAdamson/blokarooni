package config

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Data Types
type PieceData struct {
	Offset [2]int
	Cells  [][]bool
}

type PieceState struct {
	Number      int
	NumSquares  int
	IsPlaced    bool
	Orientation int
	Origin      [2]int
}

type PlayerState struct {
	Id              int
	Turn            int
	PiecesRemaining int
	Skipped         bool
	Score           int
	Pieces          [21]PieceState
}

type BoardStateEntry struct {
	PieceState
	PlayerNumber int
}

type SideboardPieceSquare struct {
	PieceNumber   int
	CollisionRect rl.Rectangle
}

type PreviewPiece struct {
	Number          int
	Origin          [2]int
	SelectionOffset [2]int
	Orientation     int
	IsVisible       bool
	Color           rl.Color
}
