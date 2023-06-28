package game

import (
	"fmt"

	c "github.com/AdamMcAdamson/blockeroni/config"

	s "github.com/AdamMcAdamson/blockeroni/state"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)
	rl.DrawText(fmt.Sprintf("Piece: %d, Orientation: %d", s.PieceToPlace, s.PieceOrientation), 10, 10, 40, rl.DarkGray)

	drawGameBoard()
	drawGameBoardPieces()
	drawPreviewBoard()
	drawPreviewBoardPiece()

	drawSkipButton()

	rl.EndDrawing()
}

func drawSkipButton() {
	text := "Skip Turn"
	var fontSize float32 = 24
	rl.DrawRectangleRounded(c.SkipTurnButtonBounds, .25, 4, rl.LightGray)
	textSize := rl.MeasureTextEx(rl.GetFontDefault(), text, fontSize, 1)
	rl.DrawText(text, c.SkipTurnButtonBounds.ToInt32().X+((c.SkipTurnButtonBounds.ToInt32().Width-int32(textSize.X))/2), c.SkipTurnButtonBounds.ToInt32().Y+((c.SkipTurnButtonBounds.ToInt32().Height-int32(textSize.Y))/2), int32(fontSize), rl.DarkGray)
}

func drawGameBoard() {
	// Draw gameBoard grid
	rl.DrawLineV(rl.Vector2Subtract(c.GameBoardStartingPos, rl.Vector2{X: 0.0, Y: 0.0}), rl.Vector2Add(c.GameBoardStartingPos, rl.Vector2{X: 0.0, Y: c.GameBoardSizePixels.Y + 1}), rl.Black)
	for i := 1; i <= c.GameBoardWidth; i++ {
		start := rl.Vector2Add(c.GameBoardStartingPos, rl.Vector2Multiply(c.CellSize, rl.Vector2{X: float32(i), Y: 0.0}))
		end := rl.Vector2Add(start, rl.Vector2{X: 0.0, Y: c.GameBoardSizePixels.Y})
		rl.DrawLineV(start, end, rl.Black)
	}
	for i := 0; i <= c.GameBoardHeight; i++ {
		start := rl.Vector2Add(c.GameBoardStartingPos, rl.Vector2Multiply(c.CellSize, rl.Vector2{X: 0.0, Y: float32(i)}))
		end := rl.Vector2Add(start, rl.Vector2{X: c.GameBoardSizePixels.X, Y: 0.0})
		rl.DrawLineV(start, end, rl.Black)
	}
}

func drawGameBoardPieces() {
	// Color grid locations
	var drawPos rl.Vector2
	for x, col := range s.GameBoard {
		for y, val := range col {
			if val != 0 {
				drawPos = rl.Vector2Add(c.GameBoardStartingPos, rl.Vector2Add(rl.Vector2Multiply(rl.Vector2{X: float32(x), Y: float32(y)}, c.CellSize), rl.Vector2{X: 0, Y: 1.0}))
				rl.DrawRectangleV(drawPos, rl.Vector2Subtract(c.CellSize, rl.Vector2{X: 1.0, Y: 1.0}), c.PlayerColor[val])
			}
		}
	}
}

func drawPreviewBoard() {
	// Draw preview grid
	rl.DrawLineV(rl.Vector2Subtract(c.PreviewBoardStartingPos, rl.Vector2{X: 1.0, Y: 1.0}), rl.Vector2Add(c.PreviewBoardStartingPos, rl.Vector2{X: 0.0, Y: c.PreviewBoardSizePixels.Y + 1}), rl.Black)
	for i := 1; i <= c.PreviewBoardWidth; i++ {
		start := rl.Vector2Add(c.PreviewBoardStartingPos, rl.Vector2Multiply(c.PreviewCellSize, rl.Vector2{X: float32(i), Y: 0.0}))
		end := rl.Vector2Add(start, rl.Vector2{X: 0.0, Y: c.PreviewBoardSizePixels.Y})
		rl.DrawLineV(start, end, rl.Black)
	}
	for i := 0; i <= c.PreviewBoardHeight; i++ {
		start := rl.Vector2Add(c.PreviewBoardStartingPos, rl.Vector2Multiply(c.PreviewCellSize, rl.Vector2{X: 0.0, Y: float32(i)}))
		end := rl.Vector2Add(start, rl.Vector2{X: c.PreviewBoardSizePixels.X, Y: 0.0})
		rl.DrawLineV(start, end, rl.Black)
	}
}

func drawPreviewBoardPiece() {
	// Color grid locations
	var drawPos rl.Vector2
	for x, col := range s.PreviewBoard {
		for y, val := range col {
			if val != 0 {
				drawPos = rl.Vector2Add(c.PreviewBoardStartingPos, rl.Vector2Add(rl.Vector2Multiply(rl.Vector2{X: float32(x), Y: float32(y)}, c.PreviewCellSize), rl.Vector2{X: 0, Y: 1.0}))
				rl.DrawRectangleV(drawPos, rl.Vector2Subtract(c.PreviewCellSize, rl.Vector2{X: 1.0, Y: 1.0}), c.PlayerColor[val])
			}
		}
	}
	drawPos = rl.Vector2Add(c.PreviewBoardStartingPos, rl.Vector2Multiply(rl.Vector2{X: 4.0, Y: 4.0}, c.PreviewCellSize))
	rl.DrawText("x", int32(drawPos.X+(c.PreviewCellWidth/4)), int32(drawPos.Y /*+(c.PreviewCellHeight/20)*/), 42, rl.Brown)
}
