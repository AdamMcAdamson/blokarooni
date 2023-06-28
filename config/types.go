package config

// Data Types
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
