package game

import (
	"fmt"

	s "github.com/AdamMcAdamson/blockeroni/state"
)

func isValidPlacementSquare(x int, y int, playerNumber int) bool {
	// Check if square exists and is free
	if x > 19 || y > 19 || x < 0 || y < 0 {
		// fmt.Printf("Invalid placementSquare, tile out of bounds. Tile (%d, %d)\n", x, y)
		return false
	}
	if s.GameBoard[x][y] != 0 {
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

func isValidPlacement(x int, y int, playerNumber int, piece int, orientation int) bool {
	// Ensure at least one square isValidConnectingSquare
	// Ensure isValidPlacementSquare for all squares

	connectionFound := false

	for py, prow := range s.Pieces[piece] {
		for px, pval := range prow {
			if pval {
				switch orientation {
				case 0:
					if !isValidPlacementSquare(x+px, y+py, playerNumber) {
						return false
					} else if !connectionFound {
						connectionFound = isValidConnectingSquare(x+px, y+py, playerNumber)
					}
				case 1:
					if !isValidPlacementSquare(x+py, y-px, playerNumber) {
						return false
					} else if !connectionFound {
						connectionFound = isValidConnectingSquare(x+py, y-px, playerNumber)
					}
				case 2:
					if !isValidPlacementSquare(x-px, y-py, playerNumber) {
						return false
					} else if !connectionFound {
						connectionFound = isValidConnectingSquare(x-px, y-py, playerNumber)
					}
				case 3:
					if !isValidPlacementSquare(x-py, y+px, playerNumber) {
						return false
					} else if !connectionFound {
						connectionFound = isValidConnectingSquare(x-py, y+px, playerNumber)
					}
				case 4:
					if !isValidPlacementSquare(x-px, y+py, playerNumber) {
						return false
					} else if !connectionFound {
						connectionFound = isValidConnectingSquare(x-px, y+py, playerNumber)
					}
				case 5:
					if !isValidPlacementSquare(x+py, y+px, playerNumber) {
						return false
					} else if !connectionFound {
						connectionFound = isValidConnectingSquare(x+py, y+px, playerNumber)
					}
				case 6:
					if !isValidPlacementSquare(x+px, y-py, playerNumber) {
						return false
					} else if !connectionFound {
						connectionFound = isValidConnectingSquare(x+px, y-py, playerNumber)
					}
				case 7:
					if !isValidPlacementSquare(x-py, y-px, playerNumber) {
						return false
					} else if !connectionFound {
						connectionFound = isValidConnectingSquare(x-py, y-px, playerNumber)
					}
				default:
					panic(fmt.Sprintf("Invalid piece orientation. Player %d, Piece %d", playerNumber, piece))
				}
			}
		}
	}
	if !connectionFound {
		fmt.Printf("No valid connecting square found.\n")
	}
	return connectionFound
}
