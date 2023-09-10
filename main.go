package main

import (
	"github.com/AdamMcAdamson/blockeroni/game"

	c "github.com/AdamMcAdamson/blockeroni/config"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// @TODO: Finish Game
// ------------------
// Load Save in Game window (Use s.GameState for saves)
// End game state and screen
// Pick starting player
// Skip Turns
// Create function that determine if players have a valid play
//  -- auto-skip if they have none (and they skipped before)
//  -- use this simulates a game from a given boardState (and randomly play games for debug)
// UI Work
// -- Clip the responsive floating piece against the game board
// -- Ensure floating pieces are the right size (may simply need int cell Width/Height)
// -- (?) Center selected pieces on Cursor
// -- Flipping feedback (for when piece looks the same even with different 'orientation') (maybe do a slight gradient on the outline?)
// -- Abstract Drawing Boards (Game and Preview) into one function
// -- General Cleanup (Say which player's turn it is, better piece highlighting etc)
// -- Buttons (rotation, skip turn, etc)
// ------------------

func main() {
	// --------------------------------------------------
	// Window Setup
	// --------------------------------------------------
	rl.InitWindow(c.WindowWidth, c.WindowHeight, "Blokarooni")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	// Init game
	game.Init()

	// Init gameBoard
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
