package config

// Data Types
type PieceState struct {
	Number      int
	IsPlaced    bool
	Orientation int
	Origin      [2]int
}

type PlayerState struct {
	Id              int
	Turn            int
	PiecesRemaining int
	Bonus           int
	Skipped         bool
	Pieces          [21]PieceState
}

type BoardStateEntry struct {
	PieceState
	PlayerNumber int
}
