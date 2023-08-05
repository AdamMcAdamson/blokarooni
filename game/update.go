package game

import (
	"fmt"

	c "github.com/AdamMcAdamson/blockeroni/config"
	rl "github.com/gen2brain/raylib-go/raylib"

	s "github.com/AdamMcAdamson/blockeroni/state"
)

func checkToEndGame() bool {
	allPlaced := true
	for i := range s.Players {
		if s.Players[i].PiecesRemaining != 0 {
			if !s.Players[i].Skipped {
				return false
			}
			allPlaced = false
		}
	}

	if allPlaced {
		s.GameState = 2 // GameOver
		return true
	} else {
		s.ShouldShowEndGameButton = true
	}
	return false
}

func skipTurn() {
	s.Players[s.CurrentPlayerIndex].Turn++
	s.Players[s.CurrentPlayerIndex].Skipped = true
	checkToEndGame()
	updateCurrentPlayer()
	updatePieceToPlace(true, false)
}

func addBoardStateEntry(playerId int, piece c.PieceState) {
	fmt.Printf("ADDING BOARD STATE ENTRY:\nPlayerId: %d\nPiece: %v\n", playerId, piece)
	s.BoardState = append(s.BoardState, c.BoardStateEntry{PieceState: piece, PlayerNumber: playerId})
}
func setGameStateAfterLoad() {
	Init()

	// Synce Player State with BoardState
	for ei, entry := range s.BoardState {
		for playerIndex := range s.Players {
			if entry.PlayerNumber == s.Players[playerIndex].Id {
				if entry.Number == 0 && s.Players[playerIndex].PiecesRemaining == 1 {
					s.Players[playerIndex].Score = 20
				} else if s.Players[playerIndex].PiecesRemaining == 1 {
					s.Players[playerIndex].Score = 15
				}

				s.Players[playerIndex].Pieces[entry.PieceState.Number] = entry.PieceState
				s.Players[playerIndex].PiecesRemaining--
				s.Players[playerIndex].Turn++
				s.Players[playerIndex].Skipped = false

				// Update Current Player and PieceToPlace
				if ei == len(s.BoardState)-1 {
					fmt.Println("SET PLAYER INDEX")
					s.CurrentPlayerIndex = playerIndex
					updateCurrentPlayer()
					updatePieceToPlace(true, false)
					s.DebugPrint()
				}
			}
		}
	}
}

// Update square if valid
func updateSquare(x int, y int, playerId int) {
	if x > 19 || y > 19 || x < 0 || y < 0 {
		fmt.Printf("Invalid boardState, tile out of bounds. Tile (%d, %d)\n", x, y)
	} else if s.GameBoard[x][y] != 0 && s.GameBoard[x][y] < 5 {
		fmt.Printf("Invalid boardState, tile conflict at (%d, %d)\n", x, y)
	} else {
		s.GameBoard[x][y] = playerId
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
		piece := s.Pieces[entry.Number]

		for iy, prow := range piece.Cells {
			for ix, pval := range prow {

				if pval {
					px := ix - piece.Offset[0]
					py := iy - piece.Offset[1]

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
	if checkToEndGame() {
		endGame()
	}
}

func updatePiecePreview(x int, y int) {
	s.PiecePreview.Number = s.PieceToPlace
	s.PiecePreview.Orientation = s.PieceOrientation
	s.PiecePreview.IsVisible = true
	s.PiecePreview.Color = rl.ColorAlpha(c.PlayerColor[s.Players[s.CurrentPlayerIndex].Id], c.PiecePreviewAlpha)
	s.PiecePreview.Origin = [2]int{x, y}
}

func updateCurrentPlayer() {
	s.PieceSelected = false
	s.PieceOrientation = 0
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

func calculateNegativeScores() {
	for i := range s.Players {
		workingScore := 0
		if s.Players[i].PiecesRemaining > 0 {
			for j := range s.Players[i].Pieces {
				if !s.Players[i].Pieces[j].IsPlaced {
					workingScore -= s.Players[i].Pieces[j].NumSquares
				}
			}
			s.Players[i].Score = workingScore
		}
	}
}

func endGame() {
	calculateNegativeScores()
}

// func clearPreviewBoard() {
// 	for x := range s.PreviewBoard {
// 		for y := range s.PreviewBoard[0] {
// 			s.PreviewBoard[x][y] = 0
// 		}
// 	}
// }

// func updatePreviewSquare(x int, y int, playerNumber int) {
// 	if x > 9 || y > 9 || x < 0 || y < 0 {
// 		fmt.Printf("Invalid previewBoard, tile out of bounds. Tile (%d, %d)\n", x, y)
// 	} else {
// 		s.PreviewBoard[x][y] = playerNumber
// 	}
// }

// func UpdatePreviewBoard(playerNumber int, piece int, orientation int) {
// 	clearPreviewBoard()
// 	x := 4
// 	y := 4
// 	for py, prow := range s.Pieces[piece] {
// 		for px, pval := range prow {
// 			// @TODO: handle player
// 			// @TODO: handle conflicts (red touching red) (Maybe we should do this on attempt to place)
// 			if pval {
// 				switch orientation {
// 				case 0:
// 					updatePreviewSquare(x+px, y+py, playerNumber)
// 				case 1:
// 					updatePreviewSquare(x+py, y-px, playerNumber)
// 				case 2:
// 					updatePreviewSquare(x-px, y-py, playerNumber)
// 				case 3:
// 					updatePreviewSquare(x-py, y+px, playerNumber)
// 				case 4:
// 					updatePreviewSquare(x-px, y+py, playerNumber)
// 				case 5:
// 					updatePreviewSquare(x+py, y+px, playerNumber)
// 				case 6:
// 					updatePreviewSquare(x+px, y-py, playerNumber)
// 				case 7:
// 					updatePreviewSquare(x-py, y-px, playerNumber)
// 				default:py :=
// 					panic(fmt.Sprintf("Invalid piece orientation. Player %d, Piece %d", playerNumber, piece))
// 				}
// 			}
// 		}
// 	}
// }
