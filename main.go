package main

import (
	"github.com/AdamMcAdamson/blockeroni/game"

	c "github.com/AdamMcAdamson/blockeroni/config"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// --------------------------------------------------
	// Window Setup
	// --------------------------------------------------
	rl.InitWindow(c.WindowWidth, c.WindowHeight, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	// Init game
	game.Init()

	// Init gameBoard
	game.UpdateBoardState()
	game.UpdateGameBoard()

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
	// Main Game loop
	// --------------------------------------------------
	for !rl.WindowShouldClose() {

		// --------------------------------------------------
		// Gameplay Logic
		// --------------------------------------------------
		game.HandleInput()

		// Clear and Update gameBoard
		game.ClearGameBoard()
		game.UpdateGameBoard()

		// game.UpdatePreviewBoard(s.Players[s.CurrentPlayerIndex].Id, s.PieceToPlace, s.PieceOrientation)

		// --------------------------------------------------
		// Drawing
		// --------------------------------------------------
		game.Draw()
	}
}
