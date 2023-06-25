package gamestate

import (
	c "github.com/AdamMcAdamson/blockeroni/config"
)

// Players Data
var Players = [4]c.PlayerState{
	{
		Id:     1,
		Pieces: [21]c.PieceState{},
	},
	{
		Id:     2,
		Pieces: [21]c.PieceState{},
	},
	{
		Id:     3,
		Pieces: [21]c.PieceState{},
	},
	{
		Id:     4,
		Pieces: [21]c.PieceState{},
	},
}
