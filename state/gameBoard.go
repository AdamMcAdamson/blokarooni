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

var GameState = 0 // 0: MainMenu, 1: Playing, 2: GameOver, 3: Paused, 4: LoadSave

const MainMenu = 0
const Playing = 1
const GameOver = 2
const Paused = 3
const LoadSave = 4

var ActiveMenuId = 0 // 0: MainMenu, 1: PauseMenu

var GameScreen rl.Texture2D

var MousePosition rl.Vector2
