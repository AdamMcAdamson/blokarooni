package game

import (
	"fmt"

	c "github.com/AdamMcAdamson/blockeroni/config"
	rl "github.com/gen2brain/raylib-go/raylib"

	s "github.com/AdamMcAdamson/blockeroni/state"
)

func StepGame() {
	switch s.GameState {
	case s.MainMenu:
		HandleInput()
		Draw()
	case s.Playing:
		HandleInput()
		clearGameBoard()
		UpdateGameBoard()
		Draw()
	case s.GameOver:
		HandleInput()
		Draw()
	case s.Paused:
		HandleInput()
		Draw()
	default:
		panic(fmt.Sprintf("StepGame(): Invalid GameState %d", s.GameState))
	}
}

func setGameState(newGameState int) {
	switch s.GameState {
	case s.MainMenu:
		setGameStateFromMainMenu(newGameState)
	case s.Playing:
		setGameStateFromPlaying(newGameState)
	case s.GameOver:
		setGameStateFromGameOver(newGameState)
	case s.Paused:
		setGameStateFromPaused(newGameState)
	}
}

func setGameStateFromPlaying(newGameState int) {
	switch newGameState {
	case s.Paused:
		s.GameScreen = rl.LoadTextureFromImage(rl.LoadImageFromScreen())
		s.ActiveMenuId = 1
		s.GameState = s.Paused
	}
}

func setGameStateFromPaused(newGameState int) {
	switch newGameState {
	case s.Playing:
		// s.GameScreen = rl.LoadTextureFromImage(rl.LoadImageFromScreen())
		s.ActiveMenuId = -1
		s.GameState = s.Playing
	}
}

func setGameStateFromMainMenu(newGameState int) {
	s.DisableAllButtons()
	switch newGameState {
	case s.MainMenu:
		s.EnableButton("MainMenuPlay")
	case s.Playing:
		clearGameBoard()
		UpdateGameBoard()
		s.GameState = s.Playing
	}
}
func setGameStateFromGameOver(newGameState int) {
	s.DisableAllButtons()
	switch newGameState {
	case s.MainMenu:
		s.EnableButton("MainMenuPlay")
		s.GameState = s.MainMenu
	case s.Playing:
		clearGameBoard()
		UpdateGameBoard()
		s.GameState = s.Playing
	}
}

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
		s.GameState = s.GameOver
		return true
	} else {
		s.ShouldShowEndGameButton = true
	}
	return false
}

// func skipTurn() {
// 	s.Players[s.CurrentPlayerIndex].Turn++
// 	s.Players[s.CurrentPlayerIndex].Skipped = true
// 	checkToEndGame()
// 	updateCurrentPlayer()
// 	updatePieceToPlace(true, false)
// }

// Append a BoardStateEntry to the BoardState
func addBoardStateEntry(playerId int, piece c.PieceState) {
	fmt.Printf("ADDING BOARD STATE ENTRY:\nPlayerId: %d\nPiece: %v\n", playerId, piece)
	s.BoardState = append(s.BoardState, c.BoardStateEntry{PieceState: piece, PlayerId: playerId})
}

func setGameStateAfterLoad() {
	Init()

	// Synce Player State with BoardState
	for ei, entry := range s.BoardState {
		for playerIndex := range s.Players {
			if entry.PlayerId == s.Players[playerIndex].Id {
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
	setGameState(1)
}

// Update the gameboard cell if valid
func updateGameBoardCell(x int, y int, playerId int) {
	if x > 19 || y > 19 || x < 0 || y < 0 {
		fmt.Printf("Invalid boardState, tile out of bounds. Tile (%d, %d)\n", x, y)
	} else if s.GameBoard[x][y] != 0 && s.GameBoard[x][y] < 5 {
		fmt.Printf("Invalid boardState, tile conflict at (%d, %d)\n", x, y)
	} else {
		s.GameBoard[x][y] = playerId
	}
}

// Sets the gameboard cells to 0 (resets the board)
func clearGameBoard() {
	for x := range s.GameBoard {
		for y := range s.GameBoard[0] {
			s.GameBoard[x][y] = 0
		}
	}
}

// Update Game Board based on boardState
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
						updateGameBoardCell(x+px, y+py, entry.PlayerId)
					case 1:
						updateGameBoardCell(x+py, y-px, entry.PlayerId)
					case 2:
						updateGameBoardCell(x-px, y-py, entry.PlayerId)
					case 3:
						updateGameBoardCell(x-py, y+px, entry.PlayerId)
					case 4:
						updateGameBoardCell(x-px, y+py, entry.PlayerId)
					case 5:
						updateGameBoardCell(x+py, y+px, entry.PlayerId)
					case 6:
						updateGameBoardCell(x+px, y-py, entry.PlayerId)
					case 7:
						updateGameBoardCell(x-py, y-px, entry.PlayerId)
					default:
						panic(fmt.Sprintf("Invalid piece orientation. Player %d, Piece %d", entry.PlayerId, entry.Number))
					}
				}
			}
		}
	}
	if checkToEndGame() {
		endGame()
	}
}

// Update game state with new piece preview
func updatePiecePreview(x int, y int) {
	s.PiecePreview.Number = s.PieceToPlace
	s.PiecePreview.Orientation = s.PieceOrientation
	s.PiecePreview.IsVisible = true
	s.PiecePreview.Color = rl.ColorAlpha(c.PlayerColor[s.Players[s.CurrentPlayerIndex].Id], c.PiecePreviewAlpha)
	s.PiecePreview.Origin = [2]int{x, y}
}

// Select the next valid player
func updateCurrentPlayer() {
	s.PieceSelected = false
	s.PieceOrientation = 0
	counter := 0

	// Wrap around
	for counter < 4 {
		if s.CurrentPlayerIndex < 3 {
			s.CurrentPlayerIndex++
		} else {
			s.CurrentPlayerIndex = 0
		}
		playerHasAValidPlacementRemaining(s.CurrentPlayerIndex)
		if s.Players[s.CurrentPlayerIndex].PiecesRemaining > 0 {
			return
		}
		counter++
	}
	fmt.Printf("ERROR: Players.PiecesRemaining and/or game over condition is not being set correctly. All players seem to have no pieces remaining.\n")
}

// Update Piece Orientation
// mode - 0: reset, 1: rotate left, 2: rotate right, 3: flip left-right, 4: flip up-down,
// 5: cycle left, 6: cycle right
func updatePieceOrientation(mode int) {
	switch mode {
	case 0:
		s.PieceOrientation = 0
	case 1:
		switch s.PieceOrientation {
		case 7:
			s.PieceOrientation = 4
		case 3:
			s.PieceOrientation = 0
		default:
			s.PieceOrientation++
		}
	case 2:
		switch s.PieceOrientation {
		case 4:
			s.PieceOrientation = 7
		case 0:
			s.PieceOrientation = 3
		default:
			s.PieceOrientation--
		}
	case 3:
		switch s.PieceOrientation {
		case 3:
			fallthrough
		case 7:
			s.PieceOrientation -= 6
		default:
			s.PieceOrientation += ((s.PieceOrientation + 1) % 2) * 4
			s.PieceOrientation += (s.PieceOrientation % 2) * 6
		}
	case 4:
		switch s.PieceOrientation {
		case 2:
			fallthrough
		case 6:
			s.PieceOrientation -= 6
		default:
			s.PieceOrientation += (s.PieceOrientation % 2) * 4
			s.PieceOrientation += ((s.PieceOrientation + 1) % 2) * 6
		}
	case 5:
		s.PieceOrientation++
	case 6:
		s.PieceOrientation--
	default:
		panic(fmt.Sprintf("Invalid 'mode' parameter %d for game.updatePieceOrientation", mode))
	}

	// wrap around (required for flip logic, and case 5/6)
	if s.PieceOrientation < 0 {
		s.PieceOrientation += 8
	} else if s.PieceOrientation > 7 {
		s.PieceOrientation -= 8
	}
}

// Update Piece to Place ::
// increaseIndex - True: Increment, False: Decrement;
// force - True: Always change, False: Don't change if you don't have to (when switching currentPlayer)
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

// Calculate how many points players should lose based on the
// total number of cells of the pieces they have not placed
func calculateNegativeScores() {
	for i := range s.Players {
		workingScore := 0
		if s.Players[i].PiecesRemaining > 0 {
			for j := range s.Players[i].Pieces {
				if !s.Players[i].Pieces[j].IsPlaced {
					workingScore -= s.Players[i].Pieces[j].NumCells
				}
			}
			s.Players[i].Score = workingScore
		}
	}
}

func endGame() {
	calculateNegativeScores()
	s.GameState = s.GameOver
}

// func clearPreviewBoard() {
// 	for x := range s.PreviewBoard {
// 		for y := range s.PreviewBoard[0] {
// 			s.PreviewBoard[x][y] = 0
// 		}
// 	}
// }

// func updatePreviewCell(x int, y int, PlayerId int) {
// 	if x > 9 || y > 9 || x < 0 || y < 0 {
// 		fmt.Printf("Invalid previewBoard, tile out of bounds. Tile (%d, %d)\n", x, y)
// 	} else {
// 		s.PreviewBoard[x][y] = PlayerId
// 	}
// }

// func UpdatePreviewBoard(PlayerId int, piece int, orientation int) {
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
// 					updatePreviewCell(x+px, y+py, PlayerId)
// 				case 1:
// 					updatePreviewCell(x+py, y-px, PlayerId)
// 				case 2:
// 					updatePreviewCell(x-px, y-py, PlayerId)
// 				case 3:
// 					updatePreviewCell(x-py, y+px, PlayerId)
// 				case 4:
// 					updatePreviewCell(x-px, y+py, PlayerId)
// 				case 5:
// 					updatePreviewCell(x+py, y+px, PlayerId)
// 				case 6:
// 					updatePreviewCell(x+px, y-py, PlayerId)
// 				case 7:
// 					updatePreviewCell(x-py, y-px, PlayerId)
// 				default:py :=
// 					panic(fmt.Sprintf("Invalid piece orientation. Player %d, Piece %d", PlayerId, piece))
// 				}
// 			}
// 		}
// 	}
// }
