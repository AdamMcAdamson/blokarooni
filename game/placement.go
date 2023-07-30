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

	checkToEndGame()

	if s.GameState != 2 {
		updateCurrentPlayer()
		updatePieceToPlace(true, false)
	}
	UpdateBoardState()
}

func isValidPlacementSquare(x int, y int, playerNumber int) bool {
	// Check if square exists and is free
	if x > 19 || y > 19 || x < 0 || y < 0 {
		// fmt.Printf("Invalid placementSquare, tile out of bounds. Tile (%d, %d)\n", x, y)
		return false
	}
	if s.GameBoard[x][y] != 0 && s.GameBoard[x][y] < 5 {
		// fmt.Printf("Invalid placementSquare, tile conflict at (%d, %d)\n", x, y)
		return false
	}

	// Ensure sides are not of same player
	if x > 0 && s.GameBoard[x-1][y] == playerNumber {
		return false
	}
	if x < 19 && s.GameBoard[x+1][y] == playerNumber {
		return false
	}
	if y > 0 && s.GameBoard[x][y-1] == playerNumber {
		return false
	}
	if y < 19 && s.GameBoard[x][y+1] == playerNumber {
		return false
	}

	return true
}

func isValidConnectingSquare(x int, y int, playerNumber int) bool {
	// Check diagonal for same playerNumber

	// Guaranteed in-bounds by isValidPlacementSquare
	// if x > 19 || y > 19 || x < 0 || y < 0 {...}

	if x > 0 {
		if y > 0 {
			if s.GameBoard[x-1][y-1] == playerNumber {
				return true
			}
		}
		if y < 19 {
			if s.GameBoard[x-1][y+1] == playerNumber {
				return true
			}
		}
	}
	if x < 19 {
		if y > 0 {
			if s.GameBoard[x+1][y-1] == playerNumber {
				return true
			}
		}
		if y < 19 {
			if s.GameBoard[x+1][y+1] == playerNumber {
				return true
			}
		}
	}
	// fmt.Printf("Invalid connectingSquare. Tile (%d, %d)\n", x, y)
	return false
}

func isCorner(x int, y int) bool {
	return (x == 0 || x == 19) && (y == 0 || y == 19)
}

func isValidPlacement(x int, y int, playerNumber int, pieceNumber int, orientation int, firstPiece bool) bool {
	// Ensure at least one square isValidConnectingSquare
	// Ensure isValidPlacementSquare for all squares

	connectionFound := false

	piece := s.Pieces[pieceNumber]

	for iy, prow := range piece.Cells {
		for ix, pval := range prow {
			if pval {
				px := ix - piece.Offset[0]
				py := iy - piece.Offset[1]

				switch orientation {
				case 0:
					if !isValidPlacementSquare(x+px, y+py, playerNumber) {
						return false
					} else if !connectionFound && !firstPiece {
						connectionFound = isValidConnectingSquare(x+px, y+py, playerNumber)
					} else if !connectionFound {
						connectionFound = isCorner(x+px, y+py)
					}
				case 1:
					if !isValidPlacementSquare(x+py, y-px, playerNumber) {
						return false
					} else if !connectionFound && !firstPiece {
						connectionFound = isValidConnectingSquare(x+py, y-px, playerNumber)
					} else if !connectionFound {
						connectionFound = isCorner(x+py, y-px)
					}
				case 2:
					if !isValidPlacementSquare(x-px, y-py, playerNumber) {
						return false
					} else if !connectionFound && !firstPiece {
						connectionFound = isValidConnectingSquare(x-px, y-py, playerNumber)
					} else if !connectionFound {
						connectionFound = isCorner(x-px, y-py)
					}
				case 3:
					if !isValidPlacementSquare(x-py, y+px, playerNumber) {
						return false
					} else if !connectionFound && !firstPiece {
						connectionFound = isValidConnectingSquare(x-py, y+px, playerNumber)
					} else if !connectionFound {
						connectionFound = isCorner(x-py, y+px)
					}
				case 4:
					if !isValidPlacementSquare(x-px, y+py, playerNumber) {
						return false
					} else if !connectionFound && !firstPiece {
						connectionFound = isValidConnectingSquare(x-px, y+py, playerNumber)
					} else if !connectionFound {
						connectionFound = isCorner(x-px, y+py)
					}
				case 5:
					if !isValidPlacementSquare(x+py, y+px, playerNumber) {
						return false
					} else if !connectionFound && !firstPiece {
						connectionFound = isValidConnectingSquare(x+py, y+px, playerNumber)
					} else if !connectionFound {
						connectionFound = isCorner(x+py, y+px)
					}
				case 6:
					if !isValidPlacementSquare(x+px, y-py, playerNumber) {
						return false
					} else if !connectionFound && !firstPiece {
						connectionFound = isValidConnectingSquare(x+px, y-py, playerNumber)
					} else if !connectionFound {
						connectionFound = isCorner(x+px, y-py)
					}
				case 7:
					if !isValidPlacementSquare(x-py, y-px, playerNumber) {
						return false
					} else if !connectionFound && !firstPiece {
						connectionFound = isValidConnectingSquare(x-py, y-px, playerNumber)
					} else if !connectionFound {
						connectionFound = isCorner(x-py, y-px)
					}
				default:
					panic(fmt.Sprintf("Invalid piece orientation. Player %d, Piece %d", playerNumber, pieceNumber))
				}
			}
		}
	}
	if !connectionFound {
		fmt.Printf("No valid connecting square found.\n")
	}
	return connectionFound
}
