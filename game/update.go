package game

import (
	"fmt"

	c "github.com/AdamMcAdamson/blockeroni/config"

	s "github.com/AdamMcAdamson/blockeroni/state"
)

func skipTurn() {
	s.Players[s.CurrentPlayerIndex].Turn++
	updateCurrentPlayerIndex()
	updatePieceToPlace(true, false)
}

// updateBoardState
// @INFO: For now this is doesn't retain move order so that we can change piece locations real time
func UpdateBoardState() {
	s.BoardState = []c.BoardStateEntry{} // @TODO: Remove when retaining move order
	for i := range s.Players {
		for j, piece := range s.Players[i].Pieces {
			if piece.IsPlaced {
				// Only add new boardStateEntries to maintain move order
				// @TODO: Uncomment when retaining move order
				/*
					found := false
					for _, stateEntry := range boardState {
						if stateEntry.number == piece.number && stateEntry.playerNumber == players[i].id {
							found = true
							break
						}
					}
					if !found {
				*/
				fmt.Printf("Piece is Placed: %d\nNumber: %d \n", j, piece.Number)
				s.BoardState = append(s.BoardState, c.BoardStateEntry{PieceState: piece, PlayerNumber: s.Players[i].Id})
				// }
			}
		}
	}
}

// Update square if valid
func updateSquare(x int, y int, playerNumber int) {
	if x > 19 || y > 19 || x < 0 || y < 0 {
		fmt.Printf("Invalid boardState, tile out of bounds. Tile (%d, %d)\n", x, y)
	} else if s.GameBoard[x][y] != 0 {
		fmt.Printf("Invalid boardState, tile conflict at (%d, %d)\n", x, y)
	} else {
		s.GameBoard[x][y] = playerNumber
	}
}

func clearPreviewBoard() {
	for x := range s.PreviewBoard {
		for y := range s.PreviewBoard[0] {
			s.PreviewBoard[x][y] = 0
		}
	}
}

func updatePreviewSquare(x int, y int, playerNumber int) {
	if x > 9 || y > 9 || x < 0 || y < 0 {
		fmt.Printf("Invalid previewBoard, tile out of bounds. Tile (%d, %d)\n", x, y)
	} else {
		s.PreviewBoard[x][y] = playerNumber
	}
}

func UpdatePreviewBoard(playerNumber int, piece int, orientation int) {
	clearPreviewBoard()
	x := 4
	y := 4
	for py, prow := range s.Pieces[piece] {
		for px, pval := range prow {
			// @TODO: handle player
			// @TODO: handle conflicts (red touching red) (Maybe we should do this on attempt to place)
			if pval {
				switch orientation {
				case 0:
					updatePreviewSquare(x+px, y+py, playerNumber)
				case 1:
					updatePreviewSquare(x+py, y-px, playerNumber)
				case 2:
					updatePreviewSquare(x-px, y-py, playerNumber)
				case 3:
					updatePreviewSquare(x-py, y+px, playerNumber)
				case 4:
					updatePreviewSquare(x-px, y+py, playerNumber)
				case 5:
					updatePreviewSquare(x+py, y+px, playerNumber)
				case 6:
					updatePreviewSquare(x+px, y-py, playerNumber)
				case 7:
					updatePreviewSquare(x-py, y-px, playerNumber)
				default:
					panic(fmt.Sprintf("Invalid piece orientation. Player %d, Piece %d", playerNumber, piece))
				}
			}
		}
	}
}

func ClearGameBoard() {
	for x := range s.GameBoard {
		for y := range s.GameBoard[0] {
			s.GameBoard[x][y] = 0
		}
	}
}

// Update gameBoard based on boardState
func UpdateGameBoard() {
	for _, entry := range s.BoardState {
		x := entry.Origin[0]
		y := entry.Origin[1]
		// fmt.Printf("Entry number: %d\n", entry.number)
		for py, prow := range s.Pieces[entry.Number] {
			for px, pval := range prow {
				// @TODO: handle player
				// @TODO: handle conflicts (red touching red) (Maybe we should do this on attempt to place)
				if pval {
					switch entry.Orientation {
					case 0:
						updateSquare(x+px, y+py, entry.PlayerNumber)
					case 1:
						updateSquare(x+py, y-px, entry.PlayerNumber)
					case 2:
						updateSquare(x-px, y-py, entry.PlayerNumber)
					case 3:
						updateSquare(x-py, y+px, entry.PlayerNumber)
					case 4:
						updateSquare(x-px, y+py, entry.PlayerNumber)
					case 5:
						updateSquare(x+py, y+px, entry.PlayerNumber)
					case 6:
						updateSquare(x+px, y-py, entry.PlayerNumber)
					case 7:
						updateSquare(x-py, y-px, entry.PlayerNumber)
					default:
						panic(fmt.Sprintf("Invalid piece orientation. Player %d, Piece %d", entry.PlayerNumber, entry.Number))
					}
				}
			}
		}
	}
}
func updateCurrentPlayerIndex() {
	counter := 0
	for counter < 4 {
		if s.CurrentPlayerIndex < 3 {
			s.CurrentPlayerIndex++
		} else {
			s.CurrentPlayerIndex = 0
		}
		if s.Players[s.CurrentPlayerIndex].PiecesRemaining > 0 {
			return
		}
		counter++
	}
	fmt.Printf("ERROR: Players.PiecesRemaining and/or game over condition is not being set correctly. All players seem to have no pieces remaining.\n")
}

func updatePieceToPlace(increaseIndex bool, force bool) {
	counter := 0
	if s.Players[s.CurrentPlayerIndex].PiecesRemaining <= 0 {
		fmt.Printf("ERROR: CurrentPlayerIndex.PiecesRemaining <= 0 when choosing piecetoPlace. CurrentPlayerIndex (%d) and/or game over condition is not being set correctly.\n", s.CurrentPlayerIndex)
	}
	for s.Players[s.CurrentPlayerIndex].Pieces[s.PieceToPlace].IsPlaced || force {
		if increaseIndex {
			if s.PieceToPlace < 20 {
				s.PieceToPlace++
			} else {
				s.PieceToPlace = 0
			}
		} else {
			if s.PieceToPlace > 0 {
				s.PieceToPlace--
			} else {
				s.PieceToPlace = 20
			}
		}
		if counter > 21 {
			fmt.Printf("ERROR: Players.PiecesRemaining and/or Pieces.IsPlaced is not being set correctly. Cycled through all pieces. CurrentPlayerIndex: %d, PiecesRemaining: %d.\n", s.CurrentPlayerIndex, s.Players[s.CurrentPlayerIndex].PiecesRemaining)
		}
		counter++
		force = false
	}
}
