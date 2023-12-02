package config

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const Eps float32 = 0.0025 // On Laptop, Eps Must be >= 0.0025 for DrawLineV

const WindowWidth = 1800
const WindowHeight = 1120

// Game Board

// Number of Cells
const GameBoardWidth int = 20
const GameBoardHeight int = 20

const GameBoardLineWidth int = 1

// Includes border
// @VERIFY: handle other window sizes well on Desktop
const GameBoardWidthPixels = 841  // windowHeight * 0.75  // 841
const GameBoardHeightPixels = 841 // GameBoardWidthPixels // 841

var GameBoardSizePixels = rl.Vector2{X: GameBoardWidthPixels, Y: GameBoardHeightPixels}

// Only includes the bottom and right-most lines border in CellWidthWithBorder
const CellWidthWithBorder float32 = float32(GameBoardWidthPixels-GameBoardLineWidth) / float32(GameBoardWidth)    // + Eps
const CellHeightWithBorder float32 = float32(GameBoardHeightPixels-GameBoardLineWidth) / float32(GameBoardHeight) // + Eps

var CellSizeWithBorder = rl.Vector2{X: float32(CellWidthWithBorder), Y: float32(CellHeightWithBorder)}

// Add Eps to ensure lines are drawn on correct side of the pixel delimiter
// @TODO: Parameterize top margin (maybe center on screen as with X)
var GameBoardStartingPos = rl.Vector2{X: float32((WindowWidth-GameBoardWidthPixels)/2) + Eps, Y: 100 + Eps}

// Accounts for upper and left-most border lines on the board
// Also Subtract 0.5:
// because DrawRectangleV rounds (0.5, 1.5) to 1.0 to draw pixel[1]
// while   DrawLineV      rounds (1.0, 2.0) to 1.5 to draw pixel[1]
var GameBoardStartingPiecePos = rl.Vector2Add(GameBoardStartingPos, rl.Vector2{X: float32(GameBoardLineWidth) - 0.5, Y: float32(GameBoardLineWidth) - 0.5})

// Preview Board (Deprecated)
/*
// @TODO: handle other preview board sizes (requires 10 to be drawn correctly atm)
// -- I suspect this is the same issue as the window size issue
// -- This should be handled once the function to draw boards is abstracted (see ../game/draw.go)
const PreviewBoardWidth int = 10
const PreviewBoardHeight int = 10

const PreviewBoardWidthPixels float32 = 420
const PreviewBoardHeightPixels float32 = 420

var PreviewBoardSizePixels = rl.Vector2{X: PreviewBoardWidthPixels, Y: PreviewBoardHeightPixels}

const PreviewCellWidth float32 = PreviewBoardWidthPixels / float32(PreviewBoardWidth)
const PreviewCellHeight float32 = PreviewBoardHeightPixels / float32(PreviewBoardHeight)

var PreviewCellSize = rl.Vector2{X: float32(PreviewCellWidth), Y: float32(PreviewCellHeight)}

var PreviewBoardStartingPos = rl.Vector2{X: 1060, Y: 100}
*/

// Define player colors
var PlayerColor = [5]rl.Color{rl.White, rl.Blue, rl.Yellow, rl.Red, rl.Green}

// var SkipTurnButtonBounds = rl.Rectangle{X: 1060.0, Y: 560.0, Width: 200.0, Height: 80.0}

const SavePath = "./saves/"
const SaveNameBase = "blokarooni-save-"

var DebugPrinted = false
