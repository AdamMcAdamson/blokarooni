package gamestate

import (
	c "github.com/AdamMcAdamson/blockeroni/config"
)

var Players = [4]c.PlayerState{}

var PieceToPlace = 0
var PieceOrientation = 0
var CurrentPlayerIndex = 0
