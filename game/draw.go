package game

import (
	"fmt"

	c "github.com/AdamMcAdamson/blockeroni/config"

	s "github.com/AdamMcAdamson/blockeroni/state"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Draw() {
	rl.BeginDrawing()

	switch s.GameState {
	case s.MainMenu:
		drawMainMenu()
	case s.Playing:
		drawPlaying()
	case s.GameOver:
		drawGameOver()
	case s.Paused:
		drawPaused()
	default:
		panic(fmt.Sprintf("Draw(): Invalid GameState %d", s.GameState))
	}

	s.DrawActiveButtons(s.MousePosition)

	c.DebugPrinted = true

	rl.EndDrawing()
}

func drawPlaying() {
	rl.ClearBackground(rl.RayWhite)
	rl.DrawText(fmt.Sprintf("PlayerIndex: %d, Piece: %d, Orientation: %d", s.CurrentPlayerIndex, s.PieceToPlace, s.PieceOrientation), 10, 10, 40, rl.DarkGray)

	drawSideboards()

	drawGameBoard()
	drawGameBoardPieces()
	drawPreviewPiece()
	drawPieceBeingHeld()
}

func drawGameOver() {
	drawEndGameScreen()
}

func drawPaused() {
	rl.ClearBackground(rl.White)
	rl.DrawTexture(s.GameScreen, 0, 0, rl.Gray)

	switch s.ActiveMenuId {
	case 1:
		drawPauseMenu()
	}
}

func drawPauseMenu() {
	var width float32 = 600.0
	var height float32 = 800.0
	menuBoxRect := rl.Rectangle{X: ((float32(c.WindowWidth) - width) / 2), Y: ((float32(c.WindowHeight) - height) / 2), Width: width, Height: height}
	rl.DrawRectangleRounded(menuBoxRect, 0.1, 1, rl.LightGray)
	rl.DrawText("Paused", menuBoxRect.ToInt32().X+140, menuBoxRect.ToInt32().Y+20, 84, rl.Black)
}

func drawMainMenu() {
	rl.ClearBackground(rl.RayWhite)

	var textWidth int32 = 670

	rl.DrawText("Blokarooni", (c.WindowWidth-textWidth)/2, 180, 128, rl.DarkGray)
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

// Draw Player Sideboards
// Sets positions of player sideboards
// @TODO: Draw sideboards from s.SideboardPieces instead of creating and drawing them seperately
// @TODO: Move to config
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
	// Vertical lines
	for i := 0; i <= c.GameBoardWidth; i++ {
		start := rl.Vector2Add(c.GameBoardStartingPos, rl.Vector2Multiply(c.CellSizeWithBorder, rl.Vector2{X: float32(i), Y: 0.0}))
		end := rl.Vector2Add(start, rl.Vector2{X: c.Eps, Y: c.GameBoardSizePixels.Y})
		size := rl.Vector2{X: float32(c.GameBoardLineWidth), Y: c.GameBoardSizePixels.Y}
		// rl.DrawRectangleV(start, size, rl.Black)
		rl.DrawLineV(start, end, rl.Black)
		if !c.DebugPrinted {
			fmt.Printf("V%d - X: %f, Y: %f; W: %f, H: %f\n", i, start.X, start.Y, size.X, size.Y)
			fmt.Printf("V%d - sX: %f, sY: %f; eX: %f, eY: %f\n", i, start.X, start.Y, end.X, end.Y)
		}
	}
	// Horizontal lines
	for i := 0; i <= c.GameBoardWidth; i++ {
		start := rl.Vector2Add(c.GameBoardStartingPos, rl.Vector2Multiply(c.CellSizeWithBorder, rl.Vector2{X: 0.0, Y: float32(i)}))
		end := rl.Vector2Add(start, rl.Vector2{X: c.GameBoardSizePixels.X, Y: c.Eps})
		// size := rl.Vector2{X: c.GameBoardSizePixels.X, Y: float32(c.GameBoardLineWidth)}
		// rl.DrawRectangleV(start, size, rl.Black)
		rl.DrawLineV(start, end, rl.Black)
		if !c.DebugPrinted {
			fmt.Printf("H%d - sX: %f, sY: %f; eX: %f, eY: %f\n", i, start.X, start.Y, end.X, end.Y)
			// fmt.Printf("V%d - X: %f, Y: %f; W: %f, H: %f\n", i, start.X, start.Y, size.X, size.Y)
		}
	}
}

// Draws the game board pieces by the given value in each cell
func drawGameBoardPieces() {
	// Color grid locations
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

func drawPieceBeingHeld() {
	if s.PiecePreview.IsVisible {
		mousePos := rl.GetMousePosition()
		drawFloatingPiece(mousePos.X, mousePos.Y, c.CellWidthWithBorder, c.CellHeightWithBorder, s.PieceToPlace, c.PlayerColor[s.Players[s.CurrentPlayerIndex].Id], rl.Black)
	}
}

// Draw Responsive Floating piece over the board, that follows the mouse
func drawFloatingPiece(x float32, y float32, cellWidth float32, cellHeight float32, pieceNumber int, color rl.Color, edgeColor rl.Color) {
	// @TODO: Clip floating cells against gameboard

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
					drawFloatingPieceCell(posX+px, posY+py, cellWidth, cellHeight, color, edgeColor)
				case 1:
					drawFloatingPieceCell(posX+py, posY-px, cellWidth, cellHeight, color, edgeColor)
				case 2:
					drawFloatingPieceCell(posX-px, posY-py, cellWidth, cellHeight, color, edgeColor)
				case 3:
					drawFloatingPieceCell(posX-py, posY+px, cellWidth, cellHeight, color, edgeColor)
				case 4:
					drawFloatingPieceCell(posX-px, posY+py, cellWidth, cellHeight, color, edgeColor)
				case 5:
					drawFloatingPieceCell(posX+py, posY+px, cellWidth, cellHeight, color, edgeColor)
				case 6:
					drawFloatingPieceCell(posX+px, posY-py, cellWidth, cellHeight, color, edgeColor)
				case 7:
					drawFloatingPieceCell(posX-py, posY-px, cellWidth, cellHeight, color, edgeColor)
				default:
					panic(fmt.Sprintf("drawFloatingPiece(): Invalid preview piece orientation %d for pieceNumber %d", s.PiecePreview.Orientation, pieceNumber))
				}
			}
		}
	}
}

func drawFloatingPieceCell(x float32, y float32, cellWidth float32, cellHeight float32, color rl.Color, edgeColor rl.Color) {
	rl.DrawRectangleRec(rl.Rectangle{X: x + 1, Y: y + 1, Width: cellWidth - 1, Height: cellHeight - 1}, color)
	rl.DrawRectangleLinesEx(rl.Rectangle{X: x, Y: y, Width: cellWidth + 1, Height: cellHeight + 1}, float32(c.GameBoardLineWidth), edgeColor)
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
						panic(fmt.Sprintf("drawPiecePreview(): Invalid preview piece orientation %d for PiecePreview", s.PiecePreview.Orientation))
					}
				}
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
