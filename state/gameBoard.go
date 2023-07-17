package gamestate

import (
	c "github.com/AdamMcAdamson/blockeroni/config"
)

var GameBoard = [c.GameBoardWidth][c.GameBoardHeight]int{}

//var PreviewBoard = [c.PreviewBoardWidth][c.PreviewBoardHeight]int{}

var SideboardPieces = [len(Players)][]c.SideboardPieceSquare{}

var BoardState = []c.BoardStateEntry{}

var ShouldShowEndGameButton = false

var GameState = 1
