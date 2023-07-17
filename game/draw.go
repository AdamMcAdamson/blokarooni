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

	drawSideboards()

	drawGameBoard()
	drawGameBoardPieces()

	// if s.GameState != 2 {
	// 	drawPreviewBoard()
	// 	drawPreviewBoardPiece()

	// 	drawSkipButton()
	// } else {
	// 	drawEndGameScreen()
	// }

	rl.EndDrawing()
}

func drawEndGameScreen() {
	endGameScreenRect := rl.Rectangle{X: 1020, Y: 100, Width: 480, Height: 340}
	rl.DrawRectangleRounded(endGameScreenRect, 0, 4, rl.LightGray)

	indexOrder := [4]int{0, 1, 2, 3}
	temp := 0
	for i := range indexOrder {
		for j := i + 1; j < 4; j++ {
			if s.Players[indexOrder[i]].Score < s.Players[indexOrder[j]].Score {
				temp = indexOrder[i]
				indexOrder[i] = indexOrder[j]
				indexOrder[j] = temp
			}
		}
	}
	for i, val := range indexOrder {
		rl.DrawText(fmt.Sprintf("Player %d Score: %d", s.Players[val].Id, s.Players[val].Score), endGameScreenRect.ToInt32().X+20, endGameScreenRect.ToInt32().Y+int32(20+i*80), 42, c.PlayerColor[s.Players[val].Id])
	}
}

// @TODO: Draw sideboards from s.SideboardPieces instead of creating and drawing them seperately
func drawSideboards() {
	drawPlayerSideboard(40, 120, 0)
	drawPlayerSideboard(1360, 120, 1)
	drawPlayerSideboard(40, 540, 2)
	drawPlayerSideboard(1360, 540, 3)
}

func drawPlayerSideboard(x int32, y int32, playerIndex int) {

	// margin of 5 in both dimensions
	drawRegionStartX := x + 20
	drawRegionStartY := y + 20

	posX := drawRegionStartX
	posY := drawRegionStartY

	sideboardRec := rl.Rectangle{X: float32(x), Y: float32(y), Width: 400, Height: 380}
	rl.DrawRectangleRounded(sideboardRec, 0.05, 1, rl.LightGray)

	for i, piece := range s.Players[playerIndex].Pieces {
		if !piece.IsPlaced {
			posX = drawRegionStartX + c.SideboardDrawOffsets[i][0]
			posY = drawRegionStartY + c.SideboardDrawOffsets[i][1]
			drawPiece(posX, posY, 20, 20, i, s.Players[playerIndex].Id)
		}
	}

}

func drawPiece(x int32, y int32, cellWidth int32, cellHeight int32, piece int, playerId int) {
	posX := x
	posY := y
	for py, prow := range s.Pieces[piece] {
		for px, pval := range prow {
			if pval {
				posX = x + (int32(px) * cellWidth)
				posY = y + (int32(py) * cellHeight)
				rl.DrawRectangleLines(posX, posY, cellWidth+1, cellHeight+1, rl.Black)
				rl.DrawRectangle(posX+1, posY+1, cellWidth-1, cellHeight-1, c.PlayerColor[playerId])
			}
		}
	}
}

// func drawSkipButton() {
// 	text := "Skip Turn"
// 	var fontSize float32 = 24
// 	rl.DrawRectangleRounded(c.SkipTurnButtonBounds, .25, 4, rl.LightGray)
// 	textSize := rl.MeasureTextEx(rl.GetFontDefault(), text, fontSize, 1)
// 	rl.DrawText(text, c.SkipTurnButtonBounds.ToInt32().X+((c.SkipTurnButtonBounds.ToInt32().Width-int32(textSize.X))/2), c.SkipTurnButtonBounds.ToInt32().Y+((c.SkipTurnButtonBounds.ToInt32().Height-int32(textSize.Y))/2), int32(fontSize), rl.DarkGray)
// }

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
				drawPos = rl.Vector2Add(c.GameBoardStartingPos, rl.Vector2Add(rl.Vector2Multiply(rl.Vector2{X: float32(x), Y: float32(y)}, c.CellSize), rl.Vector2{X: 0, Y: 0.0}))
				rl.DrawRectangleV(drawPos, rl.Vector2Subtract(c.CellSize, rl.Vector2{X: 1.0, Y: 1.0}), c.PlayerColor[val])
			}
		}
	}
}

// func drawPreviewBoard() {
// 	// Draw preview grid
// 	rl.DrawLineV(rl.Vector2Subtract(c.PreviewBoardStartingPos, rl.Vector2{X: 1.0, Y: 1.0}), rl.Vector2Add(c.PreviewBoardStartingPos, rl.Vector2{X: 0.0, Y: c.PreviewBoardSizePixels.Y + 1}), rl.Black)
// 	for i := 1; i <= c.PreviewBoardWidth; i++ {
// 		start := rl.Vector2Add(c.PreviewBoardStartingPos, rl.Vector2Multiply(c.PreviewCellSize, rl.Vector2{X: float32(i), Y: 0.0}))
// 		end := rl.Vector2Add(start, rl.Vector2{X: 0.0, Y: c.PreviewBoardSizePixels.Y})
// 		rl.DrawLineV(start, end, rl.Black)
// 	}
// 	for i := 0; i <= c.PreviewBoardHeight; i++ {
// 		start := rl.Vector2Add(c.PreviewBoardStartingPos, rl.Vector2Multiply(c.PreviewCellSize, rl.Vector2{X: 0.0, Y: float32(i)}))
// 		end := rl.Vector2Add(start, rl.Vector2{X: c.PreviewBoardSizePixels.X, Y: 0.0})
// 		rl.DrawLineV(start, end, rl.Black)
// 	}
// }

// func drawPreviewBoardPiece() {
// 	// Color grid locations
// 	var drawPos rl.Vector2
// 	for x, col := range s.PreviewBoard {
// 		for y, val := range col {
// 			if val != 0 {
// 				drawPos = rl.Vector2Add(c.PreviewBoardStartingPos, rl.Vector2Add(rl.Vector2Multiply(rl.Vector2{X: float32(x), Y: float32(y)}, c.PreviewCellSize), rl.Vector2{X: 0, Y: 1.0}))
// 				rl.DrawRectangleV(drawPos, rl.Vector2Subtract(c.PreviewCellSize, rl.Vector2{X: 1.0, Y: 1.0}), c.PlayerColor[val])
// 			}
// 		}
// 	}
// 	drawPos = rl.Vector2Add(c.PreviewBoardStartingPos, rl.Vector2Multiply(rl.Vector2{X: 4.0, Y: 4.0}, c.PreviewCellSize))
// 	rl.DrawText("x", int32(drawPos.X+(c.PreviewCellWidth/4)), int32(drawPos.Y /*+(c.PreviewCellHeight/20)*/), 42, rl.Brown)
// }
