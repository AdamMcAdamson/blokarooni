package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	// --------------------------------------------------
	// Window Setup
	// --------------------------------------------------

	const windowWidth = 1120
	const windowHeight = 1120
	rl.InitWindow(windowWidth, windowHeight, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	const gameBoardWidth int = 20
	const gameBoardHeight int = 20

	// @TODO: handle other window sizes well (as of writing we need a multiple of 20 in order to draw correctly)
	const gameBoardWidthPixels float32 = windowHeight * 0.75   // 840
	const gameBoardHeightPixels float32 = gameBoardWidthPixels // 840

	gameBoardSizePixels := rl.Vector2{X: gameBoardWidthPixels, Y: gameBoardHeightPixels}

	const cellWidth float32 = gameBoardWidthPixels / float32(gameBoardWidth)
	const cellHeight float32 = gameBoardHeightPixels / float32(gameBoardHeight)

	cellSize := rl.Vector2{X: float32(cellWidth), Y: float32(cellHeight)}

	gameBoard := [gameBoardWidth][gameBoardHeight]int{}

	gameBoardStartingPos := rl.Vector2{X: float32((windowWidth - gameBoardWidthPixels) / 2), Y: float32((windowHeight - gameBoardHeightPixels) / 2)}

	fmt.Printf("gameBoardStaringPos (v2): %f, %f\n", gameBoardStartingPos.X, gameBoardStartingPos.Y)
	fmt.Printf("cellSize (v2): %f, %f\n", cellSize.X, cellSize.Y)

	// --------------------------------------------------
	// Data Setup
	// --------------------------------------------------

	// Define player colors
	playerColor := [5]rl.Color{rl.White, rl.Red, rl.Green, rl.Yellow, rl.Blue}

	// Pieces Structure Data
	pieces := [21][][]bool{}
	pieces[0] = [][]bool{{true}}
	pieces[1] = [][]bool{{true, true}}
	pieces[2] = [][]bool{{false, true}, {true, true}}
	pieces[3] = [][]bool{{true, true, true}}
	pieces[4] = [][]bool{{true, true, true, true}}
	pieces[5] = [][]bool{{false, false, true}, {true, true, true}}
	pieces[6] = [][]bool{{true, true, false}, {false, true, true}}
	pieces[7] = [][]bool{{true, true}, {true, true}}
	pieces[8] = [][]bool{{true, true, true}, {false, true, false}}
	pieces[9] = [][]bool{{false, true, true}, {true, true, false}, {false, true, false}}  // F
	pieces[10] = [][]bool{{true}, {true}, {true}, {true}, {true}}                         // I
	pieces[11] = [][]bool{{true, false}, {true, false}, {true, false}, {true, true}}      // L
	pieces[12] = [][]bool{{false, true}, {true, true}, {true, false}, {true, false}}      // N
	pieces[13] = [][]bool{{true, true}, {true, true}, {true, false}}                      // P
	pieces[14] = [][]bool{{true, true, true}, {false, true, false}, {false, true, false}} // T
	pieces[15] = [][]bool{{true, false, true}, {true, true, true}}                        // U
	pieces[16] = [][]bool{{false, false, true}, {false, false, true}, {true, true, true}} // V
	pieces[17] = [][]bool{{false, false, true}, {false, true, true}, {true, true, false}} // W
	pieces[18] = [][]bool{{false, true, false}, {true, true, true}, {false, true, false}} // X
	pieces[19] = [][]bool{{false, true}, {true, true}, {false, true}, {false, true}}      // Y
	pieces[20] = [][]bool{{true, true, false}, {false, true, false}, {false, true, true}} // Z

	// Data Types
	type pieceState struct {
		number      int
		isPlaced    bool
		orientation int
		origin      [2]int
	}

	type playerState struct {
		id     int
		pieces [21]pieceState
	}

	type boardStateEntry struct {
		pieceState
		playerNumber int
	}

	boardState := []boardStateEntry{}

	// Players
	players := [4]playerState{
		{
			id:     1,
			pieces: [21]pieceState{},
		},
		{
			id:     2,
			pieces: [21]pieceState{},
		},
		{
			id:     3,
			pieces: [21]pieceState{},
		},
		{
			id:     4,
			pieces: [21]pieceState{},
		},
	}

	// Player Pieces Initialization
	for i := range players {
		for j := range players[i].pieces {
			players[i].pieces[j].number = j
			players[i].pieces[j].isPlaced = false
		}
	}

	// @Debug: Piece verification
	// for i := range players[0].pieces {
	// 	players[0].pieces[i].isPlaced = true
	// 	players[0].pieces[i].orientation = 0
	// 	if i < 5 {
	// 		players[0].pieces[i].origin = [2]int{(i * 4) % 20, 4 * (i / 5)}
	// 	} else if i < 15 {
	// 		players[0].pieces[i].origin = [2]int{(i * 4) % 20, 4*(i/5) - 1}
	// 	} else {
	// 		players[0].pieces[i].origin = [2]int{(i * 4) % 20, 4*(i/5) + 1}
	// 	}
	// }

	// --------------------------------------------------
	// Game State Functions
	// --------------------------------------------------

	// updateBoardState
	// @INFO: For now this is doesn't retain move order so that we can change piece locations real time
	updateBoardState := func() {
		boardState = []boardStateEntry{} // @TODO: Remove when retaining move order
		for i := range players {
			for j, piece := range players[i].pieces {
				if piece.isPlaced {
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
					fmt.Printf("Piece is Placed: %d\nNumber: %d \n", j, piece.number)
					boardState = append(boardState, boardStateEntry{piece, players[i].id})
					// }
				}
			}
		}
	}

	// Update square if valid
	updateSquare := func(x int, y int, playerNumber int) {
		if x > 19 || y > 19 || x < 0 || y < 0 {
			fmt.Printf("Invalid boardState, tile out of bounds. Tile (%d, %d)\n", x, y)
		} else if gameBoard[x][y] != 0 {
			fmt.Printf("Invalid boardState, tile conflict at (%d, %d)\n", x, y)
		} else {
			gameBoard[x][y] = playerNumber
		}
	}

	clearGameBoard := func() {
		for x := range gameBoard {
			for y := range gameBoard[0] {
				gameBoard[x][y] = 0
			}
		}
	}

	// Update gameBoard based on boardState
	updateGameBoard := func() {
		for _, entry := range boardState {
			x := entry.origin[0]
			y := entry.origin[1]
			// fmt.Printf("Entry number: %d\n", entry.number)
			for py, prow := range pieces[entry.number] {
				for px, pval := range prow {
					// @TODO: handle player
					// @TODO: handle conflicts (red touching red) (Maybe we should do this on attempt to place)
					if pval {
						switch entry.orientation {
						case 0:
							updateSquare(x+px, y+py, entry.playerNumber)
						case 1:
							updateSquare(x+py, y-px, entry.playerNumber)
						case 2:
							updateSquare(x-px, y-py, entry.playerNumber)
						case 3:
							updateSquare(x-py, y+px, entry.playerNumber)
						case 4:
							updateSquare(x-px, y+py, entry.playerNumber)
						case 5:
							updateSquare(x+py, y+px, entry.playerNumber)
						case 6:
							updateSquare(x+px, y-py, entry.playerNumber)
						case 7:
							updateSquare(x-py, y-px, entry.playerNumber)
						default:
							panic(fmt.Sprintf("Invalid piece orientation. Player %d, Piece %d", entry.playerNumber, entry.number))
						}
					}
				}
			}
		}
	}

	// init gameBoard
	updateBoardState()
	updateGameBoard()

	// Piece Placing Data
	pieceToPlace := 0
	pieceOrientation := 0

	// --------------------------------------------------
	// Main Game loop
	// --------------------------------------------------
	for !rl.WindowShouldClose() {

		// --------------------------------------------------
		// Gameplay Logic
		// --------------------------------------------------

		// Handle Input (Piece Selection at the moment)
		// @TODO: Move to Function
		checkForMoreKeys := true
		for checkForMoreKeys {
			keyCode := rl.GetKeyPressed()
			switch keyCode {
			case rl.KeyLeft:
				if pieceOrientation > 0 {
					pieceOrientation--
				} else {
					pieceOrientation = 7
				}
			case rl.KeyRight:
				if pieceOrientation < 7 {
					pieceOrientation++
				} else {
					pieceOrientation = 0
				}
			case rl.KeyDown:
				if pieceToPlace > 0 {
					pieceToPlace--
				} else {
					pieceToPlace = 20
				}
			case rl.KeyUp:
				if pieceToPlace < 20 {
					pieceToPlace++
				} else {
					pieceToPlace = 0
				}
			case 0:
				checkForMoreKeys = false
			}
		}

		// Place Pieces
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

			mousePosOnBoard := rl.Vector2Subtract(rl.GetMousePosition(), gameBoardStartingPos)

			if mousePosOnBoard.X <= gameBoardSizePixels.X && mousePosOnBoard.Y <= gameBoardSizePixels.Y && mousePosOnBoard.X >= 0 && mousePosOnBoard.Y >= 0 {
				// get cell
				cellV := rl.Vector2DivideV(mousePosOnBoard, cellSize)

				// get gameBoard Coordinates
				x := int(cellV.X)
				y := int(cellV.Y)

				// Place piece at given coordinates
				players[1].pieces[pieceToPlace].origin = [2]int{x, y}
				players[1].pieces[pieceToPlace].orientation = pieceOrientation
				players[1].pieces[pieceToPlace].isPlaced = true

				updateBoardState()
			}
		}

		// Clear and Update gameBoard
		clearGameBoard()
		updateGameBoard()

		// --------------------------------------------------
		// Drawing
		// --------------------------------------------------

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText(fmt.Sprintf("Piece: %d, Orientation: %d", pieceToPlace, pieceOrientation), 10, 10, 40, rl.DarkGray)

		// Draw grid
		rl.DrawLineV(rl.Vector2Subtract(gameBoardStartingPos, rl.Vector2{X: 0.0, Y: 1.0}), rl.Vector2Add(gameBoardStartingPos, rl.Vector2{X: 0.0, Y: gameBoardSizePixels.Y}), rl.Black)
		for i := 1; i <= gameBoardWidth; i++ {
			start := rl.Vector2Add(gameBoardStartingPos, rl.Vector2Multiply(cellSize, rl.Vector2{X: float32(i), Y: 0.0}))
			end := rl.Vector2Add(start, rl.Vector2{X: 0.0, Y: gameBoardSizePixels.Y})
			rl.DrawLineV(start, end, rl.Black)
		}
		for i := 0; i <= gameBoardHeight; i++ {
			start := rl.Vector2Add(gameBoardStartingPos, rl.Vector2Multiply(cellSize, rl.Vector2{X: 0.0, Y: float32(i)}))
			end := rl.Vector2Add(start, rl.Vector2{X: gameBoardSizePixels.X, Y: 0.0})
			rl.DrawLineV(start, end, rl.Black)
		}

		// Color grid locations
		var drawPos rl.Vector2
		for x, col := range gameBoard {
			for y, val := range col {
				if val != 0 {
					drawPos = rl.Vector2Add(gameBoardStartingPos, rl.Vector2Multiply(rl.Vector2{X: float32(x), Y: float32(y)}, cellSize))
					rl.DrawRectangleV(drawPos, rl.Vector2Subtract(cellSize, rl.Vector2{X: 1.0, Y: 1.0}), playerColor[val])
				}
			}
		}

		rl.EndDrawing()
	}
}
