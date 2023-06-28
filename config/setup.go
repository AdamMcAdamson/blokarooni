package config

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const WindowWidth = 1600
const WindowHeight = 1120

// Game Board
const GameBoardWidth int = 20
const GameBoardHeight int = 20

// @TODO: handle other window sizes well (as of writing we need a multiple of 20 in order to draw correctly)
const GameBoardWidthPixels float32 = 840  // windowHeight * 0.75   // 840
const GameBoardHeightPixels float32 = 840 // GameBoardWidthPixels // 840

var GameBoardSizePixels = rl.Vector2{X: GameBoardWidthPixels, Y: GameBoardHeightPixels}

const CellWidth float32 = GameBoardWidthPixels / float32(GameBoardWidth)
const CellHeight float32 = GameBoardHeightPixels / float32(GameBoardHeight)

var CellSize = rl.Vector2{X: float32(CellWidth), Y: float32(CellHeight)}

var GameBoardStartingPos = rl.Vector2{X: 100.0, Y: 100.0} // rl.Vector2{X: float32((windowWidth - GameBoardWidthPixels) / 2), Y: float32((windowHeight - GameBoardHeightPixels) / 2)}

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
var PlayerColor = [5]rl.Color{rl.White, rl.Red, rl.Green, rl.Yellow, rl.Blue}

var SkipTurnButtonBounds = rl.Rectangle{X: 1060.0, Y: 560.0, Width: 240.0, Height: 80.0}
