package game

import (
	"fmt"

	s "github.com/AdamMcAdamson/blockeroni/state"
)

func placePiece(x int, y int) {
	if s.PieceToPlace == 0 && s.Players[s.CurrentPlayerIndex].PiecesRemaining == 1 {
		s.Players[s.CurrentPlayerIndex].Score = 20
	} else if s.Players[s.CurrentPlayerIndex].PiecesRemaining == 1 {
		s.Players[s.CurrentPlayerIndex].Score = 15
	}

	s.Players[s.CurrentPlayerIndex].Pieces[s.PieceToPlace].Origin = [2]int{x, y}
	s.Players[s.CurrentPlayerIndex].Pieces[s.PieceToPlace].Orientation = s.PieceOrientation
	s.Players[s.CurrentPlayerIndex].Pieces[s.PieceToPlace].IsPlaced = true
	s.Players[s.CurrentPlayerIndex].PiecesRemaining--
	s.Players[s.CurrentPlayerIndex].Turn++
	s.Players[s.CurrentPlayerIndex].Skipped = false

	addBoardStateEntry(s.Players[s.CurrentPlayerIndex].Id, s.Players[s.CurrentPlayerIndex].Pieces[s.PieceToPlace])

	checkToEndGame()

	if s.GameMode != 2 {
		updateCurrentPlayer()
		updatePieceToPlace(true, false)
	}
}

func isValidPlacementCell(x int, y int, playerId int) bool {
	// Check if square exists and is free
	if x > 19 || y > 19 || x < 0 || y < 0 {
		// fmt.Printf("Invalid placementCell, tile out of bounds. Tile (%d, %d)\n", x, y)
		return false
	}
	if s.GameBoard[x][y] != 0 && s.GameBoard[x][y] < 5 {
		// fmt.Printf("Invalid placementCell, tile conflict at (%d, %d)\n", x, y)
		return false
	}

	// Ensure sides are not of same player
	if x > 0 && s.GameBoard[x-1][y] == playerId {
		return false
	}
	if x < 19 && s.GameBoard[x+1][y] == playerId {
		return false
	}
	if y > 0 && s.GameBoard[x][y-1] == playerId {
		return false
	}
	if y < 19 && s.GameBoard[x][y+1] == playerId {
		return false
	}

	return true
}

func isValidConnectingCell(x int, y int, playerId int) bool {
	// Check diagonal for same playerId

	// Guaranteed in-bounds by isValidPlacementCell
	// if x > 19 || y > 19 || x < 0 || y < 0 {...}

	if x > 0 {
		if y > 0 {
			if s.GameBoard[x-1][y-1] == playerId {
				return true
			}
		}
		if y < 19 {
			if s.GameBoard[x-1][y+1] == playerId {
				return true
			}
		}
	}
	if x < 19 {
		if y > 0 {
			if s.GameBoard[x+1][y-1] == playerId {
				return true
			}
		}
		if y < 19 {
			if s.GameBoard[x+1][y+1] == playerId {
				return true
			}
		}
	}
	// fmt.Printf("Invalid connectingCell. Tile (%d, %d)\n", x, y)
	return false
}

func isCorner(x int, y int) bool {
	return (x == 0 || x == 19) && (y == 0 || y == 19)
}

func isValidPlacement(x int, y int, playerId int, pieceNumber int, orientation int, isFirstPiece bool) bool {
	// Ensure at least one square isValidConnectingCell
	// Ensure isValidPlacementCell for all squares

	connectionFound := false

	piece := s.Pieces[pieceNumber]

	for iy, prow := range piece.Cells {
		for ix, pval := range prow {
			if pval {
				px := ix - piece.Offset[0]
				py := iy - piece.Offset[1]

				// fmt.Printf("Orientation: %d\n", orientation)
				switch orientation {
				case 0:
					if !isValidPlacementCell(x+px, y+py, playerId) {
						return false
					} else if !connectionFound && !isFirstPiece {
						connectionFound = isValidConnectingCell(x+px, y+py, playerId)
					} else if !connectionFound {
						connectionFound = isCorner(x+px, y+py)
					}
				case 1:
					if !isValidPlacementCell(x+py, y-px, playerId) {
						return false
					} else if !connectionFound && !isFirstPiece {
						connectionFound = isValidConnectingCell(x+py, y-px, playerId)
					} else if !connectionFound {
						connectionFound = isCorner(x+py, y-px)
					}
				case 2:
					if !isValidPlacementCell(x-px, y-py, playerId) {
						return false
					} else if !connectionFound && !isFirstPiece {
						connectionFound = isValidConnectingCell(x-px, y-py, playerId)
					} else if !connectionFound {
						connectionFound = isCorner(x-px, y-py)
					}
				case 3:
					if !isValidPlacementCell(x-py, y+px, playerId) {
						return false
					} else if !connectionFound && !isFirstPiece {
						connectionFound = isValidConnectingCell(x-py, y+px, playerId)
					} else if !connectionFound {
						connectionFound = isCorner(x-py, y+px)
					}
				case 4:
					if !isValidPlacementCell(x-px, y+py, playerId) {
						return false
					} else if !connectionFound && !isFirstPiece {
						connectionFound = isValidConnectingCell(x-px, y+py, playerId)
					} else if !connectionFound {
						connectionFound = isCorner(x-px, y+py)
					}
				case 5:
					if !isValidPlacementCell(x+py, y+px, playerId) {
						return false
					} else if !connectionFound && !isFirstPiece {
						connectionFound = isValidConnectingCell(x+py, y+px, playerId)
					} else if !connectionFound {
						connectionFound = isCorner(x+py, y+px)
					}
				case 6:
					if !isValidPlacementCell(x+px, y-py, playerId) {
						return false
					} else if !connectionFound && !isFirstPiece {
						connectionFound = isValidConnectingCell(x+px, y-py, playerId)
					} else if !connectionFound {
						connectionFound = isCorner(x+px, y-py)
					}
				case 7:
					if !isValidPlacementCell(x-py, y-px, playerId) {
						return false
					} else if !connectionFound && !isFirstPiece {
						connectionFound = isValidConnectingCell(x-py, y-px, playerId)
					} else if !connectionFound {
						connectionFound = isCorner(x-py, y-px)
					}
				default:
					panic(fmt.Sprintf("Invalid piece orientation. Player %d, Piece %d", playerId, pieceNumber))
				}
			}
		}
	}
	// if !connectionFound {
	// 	// fmt.Printf("No valid connecting square found.\n")
	// }
	return connectionFound
}

// Check if the given player has a valid placement remaining
func playerHasAValidPlacementRemaining(playerIndex int) bool {
	isFirstPiece := (s.Players[s.CurrentPlayerIndex].PiecesRemaining == 21)
	for pieceIndex := range s.Pieces {
		if !s.Players[playerIndex].Pieces[pieceIndex].IsPlaced {
			for orientation := 0; orientation < 8; orientation++ {
				for y, row := range s.GameBoard {
					for x := range row {
						if isValidPlacement(x, y, s.Players[playerIndex].Id, pieceIndex, orientation, isFirstPiece) {
							fmt.Printf("Player %d can place piece %d at (%d,%d) with orientation %d\n", s.Players[playerIndex].Id, pieceIndex, x, y, orientation)
							return true
						}
					}
				}
			}
		}
	}
	fmt.Printf("Player %d has no valid placements remaining.\n", s.Players[playerIndex].Id)
	return false
}

// @TODO: implement
// func performRandomPlacement(playerIndex int) bool {
// 	if playerHasAValidPlacementRemaining(playerIndex) {
// 		// place random piece
// 	}
// 	return false
// }
