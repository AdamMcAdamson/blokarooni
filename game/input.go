package game

import (
	c "github.com/AdamMcAdamson/blockeroni/config"

	s "github.com/AdamMcAdamson/blockeroni/state"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func HandleInput() {
	handleKeys()
	handleClicks()
}

func handleKeys() {
	checkForMoreKeys := true
	for checkForMoreKeys {
		keyCode := rl.GetKeyPressed()
		switch keyCode {
		case rl.KeyLeft:
			if s.PieceOrientation > 0 {
				s.PieceOrientation--
			} else {
				s.PieceOrientation = 7
			}
		case rl.KeyRight:
			if s.PieceOrientation < 7 {
				s.PieceOrientation++
			} else {
				s.PieceOrientation = 0
			}
		case rl.KeyDown:
			updatePieceToPlace(false, true)
		case rl.KeyUp:
			updatePieceToPlace(true, true)
		case 0:
			checkForMoreKeys = false
		}
	}
}

func handleClicks() {
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

		mousePos := rl.GetMousePosition()
		mousePosOnBoard := rl.Vector2Subtract(mousePos, c.GameBoardStartingPos)

		if rl.CheckCollisionPointRec(mousePos, c.SkipTurnButtonBounds) {
			skipTurn()
		}

		// @TODO: Create rect for gameBoard and use CheckCollisionPointRec to check bounds
		if mousePosOnBoard.X <= c.GameBoardSizePixels.X && mousePosOnBoard.Y <= c.GameBoardSizePixels.Y && mousePosOnBoard.X >= 0 && mousePosOnBoard.Y >= 0 {
			// get cell
			cellV := rl.Vector2DivideV(mousePosOnBoard, c.CellSize)

			// get gameBoard Coordinates
			x := int(cellV.X)
			y := int(cellV.Y)

			if !s.Players[s.CurrentPlayerIndex].Pieces[s.PieceToPlace].IsPlaced {
				if isValidPlacement(x, y, s.Players[s.CurrentPlayerIndex].Id, s.PieceToPlace, s.PieceOrientation, (s.Players[s.CurrentPlayerIndex].PiecesRemaining == 21)) {
					placePiece(x, y)
				}
			}
		}
	}
}
