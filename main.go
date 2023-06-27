package main

import (
	"github.com/AdamMcAdamson/blockeroni/game"

	c "github.com/AdamMcAdamson/blockeroni/config"

	s "github.com/AdamMcAdamson/blockeroni/state"

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

	// @DEBUG: Setup First Piece for testing
	// s.Players[1].Pieces[0].Origin = [2]int{0, 0}
	// s.Players[1].Pieces[0].Orientation = 0
	// s.Players[1].Pieces[0].IsPlaced = true

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

		game.UpdatePreviewBoard(s.Players[s.CurrentPlayerIndex].Id, s.PieceToPlace, s.PieceOrientation)

		// --------------------------------------------------
		// Drawing
		// --------------------------------------------------
		game.Draw()
	}
}
