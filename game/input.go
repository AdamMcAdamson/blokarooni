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
			if s.PieceToPlace > 0 {
				s.PieceToPlace--
			} else {
				s.PieceToPlace = 20
			}
		case rl.KeyUp:
			if s.PieceToPlace < 20 {
				s.PieceToPlace++
			} else {
				s.PieceToPlace = 0
			}
		case 0:
			checkForMoreKeys = false
		}
	}
}

func handleClicks() {
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

		mousePosOnBoard := rl.Vector2Subtract(rl.GetMousePosition(), c.GameBoardStartingPos)

		if mousePosOnBoard.X <= c.GameBoardSizePixels.X && mousePosOnBoard.Y <= c.GameBoardSizePixels.Y && mousePosOnBoard.X >= 0 && mousePosOnBoard.Y >= 0 {
			// get cell
			cellV := rl.Vector2DivideV(mousePosOnBoard, c.CellSize)

			// get gameBoard Coordinates
			x := int(cellV.X)
			y := int(cellV.Y)

			if !s.Players[1].Pieces[s.PieceToPlace].IsPlaced {
				if isValidPlacement(x, y, s.Players[1].Id, s.PieceToPlace, s.PieceOrientation) {
					s.Players[1].Pieces[s.PieceToPlace].Origin = [2]int{x, y}
					s.Players[1].Pieces[s.PieceToPlace].Orientation = s.PieceOrientation
					s.Players[1].Pieces[s.PieceToPlace].IsPlaced = true
					if s.PieceToPlace < 20 {
						s.PieceToPlace++
					} else {
						s.PieceToPlace = 1
					}
					UpdateBoardState()
				}
			}
		}
	}
}
