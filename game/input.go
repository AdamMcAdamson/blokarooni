package game

import (
	c "github.com/AdamMcAdamson/blockeroni/config"

	s "github.com/AdamMcAdamson/blockeroni/state"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func HandleInput() {
	handleKeys()
	handleClicks()
	handleMousePositionOnBoard()
}

func handleKeys() {
	checkForMoreKeys := true
	for checkForMoreKeys {
		keyCode := rl.GetKeyPressed()
		switch keyCode {
		case rl.KeyN:
			{
				saveGame()
			}
		case rl.KeyL:
			{
				if getSaveFiles() {
					setGameStateAfterLoad()
				}
			}
		case rl.KeyA:
			updatePieceOrientation(1)
		case rl.KeyD:
			updatePieceOrientation(2)
		case rl.KeyE:
			fallthrough
		case rl.KeyQ:
			updatePieceOrientation(3)
		case rl.KeyW:
			fallthrough
		case rl.KeyS:
			updatePieceOrientation(4)
		case rl.KeyLeft:
			updatePieceOrientation(5)
		case rl.KeyRight:
			updatePieceOrientation(6)
		case rl.KeyDown:
			fallthrough
		case rl.KeyZ:
			updatePieceToPlace(false, true)
		case rl.KeyUp:
			fallthrough
		case rl.KeyX:
			updatePieceToPlace(true, true)
		case rl.KeyEscape:
			switch s.GameState {
			case 1:
				setGameState(3)
			case 3:
				setGameState(1)
			}
		case 0:
			checkForMoreKeys = false
		}
	}
}

func handleClicks() {
	s.MousePosition = rl.GetMousePosition()
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		mousePosOnBoard := rl.Vector2Subtract(s.MousePosition, c.GameBoardStartingPos)
		btnDown, _ := s.DetectAndHandleButtonDown(s.MousePosition)
		if btnDown {
			return
		}
		// @TODO: Create rect for gameBoard and use CheckCollisionPointRec to check bounds
		if s.PieceSelected && mousePosOnBoard.X <= c.GameBoardSizePixels.X && mousePosOnBoard.Y <= c.GameBoardSizePixels.Y && mousePosOnBoard.X >= 0 && mousePosOnBoard.Y >= 0 {
			// get cell
			cellV := rl.Vector2DivideV(mousePosOnBoard, c.CellSizeWithBorder)

			// get gameBoard Coordinates
			x := int(cellV.X)
			y := int(cellV.Y)

			if !s.Players[s.CurrentPlayerIndex].Pieces[s.PieceToPlace].IsPlaced {
				if isValidPlacement(x, y, s.Players[s.CurrentPlayerIndex].Id, s.PieceToPlace, s.PieceOrientation, (s.Players[s.CurrentPlayerIndex].PiecesRemaining == 21)) {
					placePiece(x, y)
				}
			}
		} else {
			for i := range s.SideboardPieces[s.CurrentPlayerIndex] {
				if !s.Players[s.CurrentPlayerIndex].Pieces[s.SideboardPieces[s.CurrentPlayerIndex][i].PieceNumber].IsPlaced && rl.CheckCollisionPointRec(s.MousePosition, s.SideboardPieces[s.CurrentPlayerIndex][i].CollisionRect) && (!s.PieceSelected || s.PieceToPlace != s.SideboardPieces[s.CurrentPlayerIndex][i].PieceNumber) {
					s.PieceToPlace = s.SideboardPieces[s.CurrentPlayerIndex][i].PieceNumber
					s.PieceSelected = true
					s.PieceOrientation = 0
				}
			}
		}
	} else if rl.IsMouseButtonPressed(rl.MouseRightButton) {
		s.PieceSelected = false
		s.PieceOrientation = 0
	}

	if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		// btnPressed, _ := s.DetectAndHandleButtonRelease(s.MousePosition)
		s.DetectAndHandleButtonRelease(s.MousePosition)
	}
}

func handleMousePositionOnBoard() {
	mousePosOnBoard := rl.Vector2Subtract(s.MousePosition, c.GameBoardStartingPos)

	if s.PieceSelected {
		// @TODO: Create rect for gameBoard and use CheckCollisionPointRec to check bounds
		if mousePosOnBoard.X <= c.GameBoardSizePixels.X && mousePosOnBoard.Y <= c.GameBoardSizePixels.Y && mousePosOnBoard.X >= 0 && mousePosOnBoard.Y >= 0 {
			//fmt.Printf("Piece Selected, mouse on Board\n")
			// get cell
			cellV := rl.Vector2DivideV(mousePosOnBoard, c.CellSizeWithBorder)

			// get gameBoard Coordinates
			x := int(cellV.X)
			y := int(cellV.Y)

			updatePiecePreview(x, y)

		} else {
			s.PiecePreview.IsVisible = false
		}
	} else {
		s.PiecePreview.IsVisible = false
	}
}
