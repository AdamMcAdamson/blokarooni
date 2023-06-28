package gamestate

import (
	c "github.com/AdamMcAdamson/blockeroni/config"
)

func Init() {
	for i := range Players {
		Players[i].Id = i + 1
		Players[i].Turn = 0
		Players[i].PiecesRemaining = 21
		Players[i].Score = 0
		Players[i].Skipped = false
		Players[i].Pieces = [21]c.PieceState{}
		for j := range Players[i].Pieces {
			Players[i].Pieces[j].Number = j
			Players[i].Pieces[j].IsPlaced = false
			Players[i].Pieces[j].NumSquares = getPieceNumSquaresFromIndex(j)
		}
	}
}
