package gamestate

import (
	c "github.com/AdamMcAdamson/blockeroni/config"
)

// Players Data
var Players = [4]c.PlayerState{
	{
		Id:              1,
		Turn:            0,
		PiecesRemaining: 21,
		Pieces:          [21]c.PieceState{},
	},
	{
		Id:              2,
		Turn:            0,
		PiecesRemaining: 21,
		Pieces:          [21]c.PieceState{},
	},
	{
		Id:              3,
		Turn:            0,
		PiecesRemaining: 21,
		Pieces:          [21]c.PieceState{},
	},
	{
		Id:              4,
		Turn:            0,
		PiecesRemaining: 21,
		Pieces:          [21]c.PieceState{},
	},
}

var PieceToPlace = 1
var PieceOrientation = 0
var CurrentPlayerIndex = 0
