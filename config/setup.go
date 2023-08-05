package config

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const eps float32 = 0.00001

const WindowWidth = 1800
const WindowHeight = 1120

// Game Board
const GameBoardWidth int = 20
const GameBoardHeight int = 20

// @TODO: handle other window sizes well (as of writing we need a multiple of 20 in order to draw correctly)
const GameBoardWidthPixels = 840  // windowHeight * 0.75   // 840
const GameBoardHeightPixels = 840 // GameBoardWidthPixels // 840

var GameBoardSizePixels = rl.Vector2{X: GameBoardWidthPixels, Y: GameBoardHeightPixels}

const CellWidth float32 = GameBoardWidthPixels/float32(GameBoardWidth) + eps
const CellHeight float32 = GameBoardHeightPixels/float32(GameBoardHeight) + eps

var CellSize = rl.Vector2{X: float32(CellWidth), Y: float32(CellHeight)}

var GameBoardStartingPos = rl.Vector2{X: float32((WindowWidth - GameBoardWidthPixels) / 2), Y: 100 /*float32((WindowHeight - GameBoardHeightPixels) / 2)*/}

// = rl.Vector2{X: 100.0, Y: 100.0}

// Preview Board
// @TODO: handle other preview board sizes (requires 10 to be drawn correctly atm)
//
//	-- I suspect this is the same issue as the window size issue
const PreviewBoardWidth int = 10
const PreviewBoardHeight int = 10

const PreviewBoardWidthPixels float32 = 420
const PreviewBoardHeightPixels float32 = 420

var PreviewBoardSizePixels = rl.Vector2{X: PreviewBoardWidthPixels, Y: PreviewBoardHeightPixels}

const PreviewCellWidth float32 = PreviewBoardWidthPixels / float32(PreviewBoardWidth)
const PreviewCellHeight float32 = PreviewBoardHeightPixels / float32(PreviewBoardHeight)

var PreviewCellSize = rl.Vector2{X: float32(PreviewCellWidth), Y: float32(PreviewCellHeight)}

var PreviewBoardStartingPos = rl.Vector2{X: 1060, Y: 100}

// Define player colors
var PlayerColor = [5]rl.Color{rl.White, rl.Blue, rl.Yellow, rl.Red, rl.Green}

// var SkipTurnButtonBounds = rl.Rectangle{X: 1060.0, Y: 560.0, Width: 200.0, Height: 80.0}

const SaveFilePath = "./saves/"
const SaveFileNameBase = "blokarooni-save-"
