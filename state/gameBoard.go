package gamestate

import (
	c "github.com/AdamMcAdamson/blockeroni/config"
)

var GameBoard = [c.GameBoardWidth][c.GameBoardHeight]int{}

//var PreviewBoard = [c.PreviewBoardWidth][c.PreviewBoardHeight]int{}

var PiecePreview = c.PreviewPiece{}

var SideboardPieces = [len(Players)][]c.SideboardPieceCell{}

var BoardState = []c.BoardStateEntry{}

var ShouldShowEndGameButton = false

var GameMode = -1
