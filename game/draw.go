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
	rl.DrawText(fmt.Sprintf("PlayerIndex: %d, Piece: %d, Orientation: %d", s.CurrentPlayerIndex, s.PieceToPlace, s.PieceOrientation), 10, 10, 40, rl.DarkGray)

	drawSideboards()

	drawGameBoard()
	drawGameBoardPieces()
	drawPreviewPiece()
	// drawPieceBeingHeld()

	// if s.GameState != 2 {
	// 	drawPreviewBoard()
	// 	drawPreviewBoardPiece()

	// 	drawSkipButton()
	// } else {
	// 	drawEndGameScreen()
	// }

	// drawCellDebug(0, 0)
	// drawCellDebug(1, 1)

	// @TODO: verify works on Desktop
	// @DebugRemove

	// drawLineVDebug(483.002502, 100.002502, 483.005005, 934.002502)
	// drawLineVDebug(524.652527, 100.002502, 524.655029, 934.002502)
	// drawLineVDebug(566.302490, 100.002502, 566.304993, 934.002502)
	// drawLineVDebug(483.002502, 100.002502, 1317.002441, 100.005005)
	// drawLineVDebug(483.002502, 141.652496, 1317.002441, 141.654999)
	// drawLineVDebug(483.002502, 183.302505, 1317.002441, 183.305008)
	// drawRectVDebugHelper(484.450, 101.40, 40, 40)  // rounds to top left of (484, 101)
	// drawRectVDebugHelper(524.601, 141.601, 41, 41) // rounds to top left of (525, 142)
	// drawPixelsForDebug()
	// drawRectVDebug()

	// -----------
	c.DebugPrinted = true

	rl.EndDrawing()
}

/*
// @DebugRemove
func drawLineVDebug(sx float32, sy float32, ex float32, ey float32) {
	rl.DrawLineV(rl.Vector2{X: sx, Y: sy}, rl.Vector2{X: ex, Y: ey}, rl.Black)
}

// @DebugRemove
func drawPixelsForDebug() {

	// top left of Draw Rect Region
	rl.DrawPixel(484, 101, rl.Red)
	rl.DrawPixel(525, 142, rl.Red)

	rl.DrawPixel(483, 250, rl.Red)
	rl.DrawPixel(483, 251, rl.Red)
	rl.DrawPixel(483, 252, rl.Red)

	rl.DrawPixel(524, 250, rl.Red)
	rl.DrawPixel(524, 251, rl.Red)
	rl.DrawPixel(524, 252, rl.Red)
	// rl.DrawPixel(1272, 250, rl.Red)
	// rl.DrawPixel(1272, 251, rl.Red)
	// rl.DrawPixel(1272, 252, rl.Red)

}

// @DebugRemove
func drawRectVDebug() {
	drawRectVDebugHelper(804.75, 9.75, 1, 1)
	rl.DrawPixel(805, 10, rl.Red)
}

// @DebugRemove
func drawRectVDebugHelper(x float32, y float32, sizeX float32, sizeY float32) {
	drawPos := rl.Vector2{X: x, Y: y}
	size := rl.Vector2{X: sizeX, Y: sizeY}
	rl.DrawRectangleV(drawPos, size, rl.Blue)
}

// @DebugRemove
func drawCellDebug(x int, y int) {
	drawPos := rl.Vector2Add(c.GameBoardStartingPiecePos, rl.Vector2Multiply(rl.Vector2{X: float32(x), Y: float32(y)}, c.CellSizeWithBorder))
	size := rl.Vector2Subtract(c.CellSizeWithBorder, rl.Vector2{X: float32(c.GameBoardLineWidth), Y: float32(c.GameBoardLineWidth)})
	rl.DrawRectangleV(drawPos, size, rl.Green)
	if !c.DebugPrinted {
		fmt.Printf("Cell (%d, %d) - X: %f, Y: %f; Width: %f, Length: %f\n", x, y, drawPos.X, drawPos.Y, size.X, size.Y)
	}
}
*/

/*
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
*/

// Draw Player Sideboards
// Sets positions of player sideboards
// @TODO: Draw sideboards from s.SideboardPieces instead of creating and drawing them seperately
func drawSideboards() {
	drawPlayerSideboard(40, 120, 0)
	drawPlayerSideboard(1360, 120, 1)
	drawPlayerSideboard(40, 540, 2)
	drawPlayerSideboard(1360, 540, 3)
}

// Draws a Player Sideboard given the location and player
// These show the players which pieces are still available to place
func drawPlayerSideboard(x int32, y int32, playerIndex int) {
	// margin in both dimensions
	drawRegionStartX := x + 20
	drawRegionStartY := y + 20

	posX := drawRegionStartX
	posY := drawRegionStartY

	// Draw the actual "board"
	sideboardRec := rl.Rectangle{X: float32(x), Y: float32(y), Width: 400, Height: 380}
	rl.DrawRectangleRounded(sideboardRec, 0.05, 1, rl.LightGray)

	// Draw the player's remaining pieces
	for i, piece := range s.Players[playerIndex].Pieces {
		if !piece.IsPlaced {
			posX = drawRegionStartX + c.SideboardDrawOffsets[i][0]
			posY = drawRegionStartY + c.SideboardDrawOffsets[i][1]
			drawSideboardPiece(posX, posY, 20, 20, i, s.Players[playerIndex].Id)
		}
	}

}

// Draws a Sideboard Piece at the given location
func drawSideboardPiece(x int32, y int32, cellWidth int32, cellHeight int32, pieceNumber int, playerId int) {
	posX := x
	posY := y
	piece := s.Pieces[pieceNumber]
	for py, prow := range piece.Cells {
		for px, pval := range prow {
			// Only draw cells that exist
			if pval {
				// Since we are drawing on a per-cell basis, we need the location of the current cell
				posX = x + (int32(px) * cellWidth)
				posY = y + (int32(py) * cellHeight)
				color := c.PlayerColor[playerId]

				// Draw an outline around the cell of the piece
				rl.DrawRectangleLines(posX, posY, cellWidth+1, cellHeight+1, rl.Black)

				// If the piece is selected by the player whose turn it is
				// fill the cell with a gradient, otherwise just use the normal player color
				if s.Players[s.CurrentPlayerIndex].Id == playerId && s.PieceToPlace == pieceNumber {
					rect := rl.Rectangle{X: float32(posX + 1), Y: float32(posY + 1), Width: float32(cellWidth - 1), Height: float32(cellHeight - 1)}
					rl.DrawRectangleGradientEx(rect, color, rl.LightGray, color, color)
				} else {
					rl.DrawRectangle(posX+1, posY+1, cellWidth-1, cellHeight-1, color)
				}
			}
		}
	}
}

// @TODO: Abstract to handle alternative sizing (to enable preview board to use the same function)
func drawGameBoard() {
	// fmt.Printf("X: %f, Y: %f\n", c.GameBoardStartingPos.X, c.GameBoardStartingPos.Y)
	// Draw gameBoard grid
	// @VERIFY: On Desktop
	for i := 0; i <= c.GameBoardWidth; i++ {
		start := rl.Vector2Add(c.GameBoardStartingPos, rl.Vector2Multiply(c.CellSizeWithBorder, rl.Vector2{X: float32(i), Y: 0.0}))
		end := rl.Vector2Add(start, rl.Vector2{X: c.Eps, Y: c.GameBoardSizePixels.Y})
		rl.DrawLineV(start, end, rl.Black)
		if !c.DebugPrinted {
			fmt.Printf("V%d - sX: %f, sY: %f; eX: %f, eY: %f\n", i, start.X, start.Y, end.X, end.Y)
		}
	}
	for i := 0; i <= c.GameBoardWidth; i++ {
		start := rl.Vector2Add(c.GameBoardStartingPos, rl.Vector2Multiply(c.CellSizeWithBorder, rl.Vector2{X: 0.0, Y: float32(i)}))
		end := rl.Vector2Add(start, rl.Vector2{X: c.GameBoardSizePixels.X, Y: c.Eps})
		rl.DrawLineV(start, end, rl.Black)
		if !c.DebugPrinted {
			fmt.Printf("H%d - sX: %f, sY: %f; eX: %f, eY: %f\n", i, start.X, start.Y, end.X, end.Y)
		}
	}
	/*
		// Old Broken Code
		rl.DrawLineV(rl.Vector2Subtract(c.GameBoardStartingPos, rl.Vector2{X: 0.0, Y: 1.0}), rl.Vector2Add(c.GameBoardStartingPos, rl.Vector2{X: 0.0, Y: c.GameBoardSizePixels.Y}), rl.Black)
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
	*/
}

// Draws the game board pieces by the given value in each cell
func drawGameBoardPieces() {
	// Color grid locations
	//var drawPos rl.Vector2
	for x, col := range s.GameBoard {
		for y, val := range col {
			if val != 0 {
				drawGameBoardCellColor(x, y, c.PlayerColor[val])
			}
		}
	}
}

// @TODO: Abstract to handle alternative sizing (to enable preview board to use the same function)
func drawGameBoardCellColor(x int, y int, color rl.Color) {
	drawPos := rl.Vector2Add(c.GameBoardStartingPiecePos, rl.Vector2Multiply(rl.Vector2{X: float32(x), Y: float32(y)}, c.CellSizeWithBorder))
	size := rl.Vector2Subtract(c.CellSizeWithBorder, rl.Vector2{X: float32(c.GameBoardLineWidth), Y: float32(c.GameBoardLineWidth)})
	rl.DrawRectangleV(drawPos, size, color)
	if !c.DebugPrinted {
		fmt.Printf("Cell (%d, %d) - X: %f, Y: %f; Width: %f, Length: %f\n", x, y, drawPos.X, drawPos.Y, size.X, size.Y)
	}
}

func drawGameBoardCellColorIfValid(x int, y int, color rl.Color) {
	if x >= 0 && x < c.GameBoardWidth && y >= 0 && y < c.GameBoardHeight {
		// fmt.Printf("Cell is valid for drawCellColor. x: %d, y: %d, Color: (%d,%d,%d,%d)\n", x, y, color.R, color.G, color.B, color.A)
		drawGameBoardCellColor(x, y, color)
	}
}

// Draw Responsive Floating piece over the board, that follows the mouse
func drawFloatingPiece(x float32, y float32, cellWidth float32, cellHeight float32, pieceNumber int) {
	// @TODO: Clip floating cells against gameboard
	color := rl.Black
	piece := s.Pieces[pieceNumber]

	for iy, prow := range piece.Cells {
		for ix, pval := range prow {
			// Only draw cells that exist
			if pval {

				// Since we are drawing on a per-cell basis, we need the location of the current cell
				posX := x - (cellWidth / 2)
				posY := y - (cellHeight / 2)

				px := float32(ix-piece.Offset[0]) * cellWidth
				py := float32(iy-piece.Offset[1]) * cellHeight

				switch s.PiecePreview.Orientation {
				case 0:
					drawFloatingPieceCell(posX+px, posY+py, cellWidth, cellHeight, color)
				case 1:
					drawFloatingPieceCell(posX+py, posY-px, cellWidth, cellHeight, color)
				case 2:
					drawFloatingPieceCell(posX-px, posY-py, cellWidth, cellHeight, color)
				case 3:
					drawFloatingPieceCell(posX-py, posY+px, cellWidth, cellHeight, color)
				case 4:
					drawFloatingPieceCell(posX-px, posY+py, cellWidth, cellHeight, color)
				case 5:
					drawFloatingPieceCell(posX+py, posY+px, cellWidth, cellHeight, color)
				case 6:
					drawFloatingPieceCell(posX+px, posY-py, cellWidth, cellHeight, color)
				case 7:
					drawFloatingPieceCell(posX-py, posY-px, cellWidth, cellHeight, color)
				default:
					panic("Invalid piece orientation!")
				}
			}
		}
	}
}

func drawFloatingPieceCell(x float32, y float32, cellWidth float32, cellHeight float32, color rl.Color) {
	rl.DrawRectangleLinesEx(rl.Rectangle{X: x, Y: y, Width: cellWidth + 1, Height: cellHeight + 1}, float32(c.GameBoardLineWidth), color)
}

// Draws an opaque piece where the a piece would be placed
// if the player were to place a piece right then
func drawPreviewPiece() {
	if s.PiecePreview.IsVisible {

		x := s.PiecePreview.Origin[0]
		y := s.PiecePreview.Origin[1]

		// Debug
		// rl.DrawText(fmt.Sprintf("X: %d, Y: %d", x, y), 10, 50, 40, rl.DarkGray)

		color := s.PiecePreview.Color

		piece := s.Pieces[s.PiecePreview.Number]

		for iy, prow := range piece.Cells {
			for ix, pval := range prow {
				if pval {
					px := ix - piece.Offset[0]
					py := iy - piece.Offset[1]

					switch s.PiecePreview.Orientation {
					case 0:
						drawGameBoardCellColorIfValid(x+px, y+py, color)
					case 1:
						drawGameBoardCellColorIfValid(x+py, y-px, color)
					case 2:
						drawGameBoardCellColorIfValid(x-px, y-py, color)
					case 3:
						drawGameBoardCellColorIfValid(x-py, y+px, color)
					case 4:
						drawGameBoardCellColorIfValid(x-px, y+py, color)
					case 5:
						drawGameBoardCellColorIfValid(x+py, y+px, color)
					case 6:
						drawGameBoardCellColorIfValid(x+px, y-py, color)
					case 7:
						drawGameBoardCellColorIfValid(x-py, y-px, color)
					default:
						panic("Invalid preview piece orientation.")
					}
				}
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
