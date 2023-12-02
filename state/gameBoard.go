package gamestate

import (
	c "github.com/AdamMcAdamson/blockeroni/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var GameBoard = [c.GameBoardWidth][c.GameBoardHeight]int{}

//var PreviewBoard = [c.PreviewBoardWidth][c.PreviewBoardHeight]int{}

var PiecePreview = c.PreviewPiece{}

var SideboardPieces = [len(Players)][]c.SideboardPieceCell{}

var BoardState = []c.BoardStateEntry{}

var ShouldShowEndGameButton = false

var GameMode = 0 // 0: MainMenu, 1: Playing, 2: GameOver, 3: Paused, 4: LoadSave

var MousePosition rl.Vector2
