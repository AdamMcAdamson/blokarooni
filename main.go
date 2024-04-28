package main

import (
	"github.com/AdamMcAdamson/blockeroni/game"

	c "github.com/AdamMcAdamson/blockeroni/config"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// @TODO: Finish Game
// ------------------
// Refactor game functionality to be based on GameState
// -- Add menus (s.ActiveMenuId) for GameStates that have multiple menus (paused)
// -- Draw
// -- Buttons (move from state.buttons to game.input/buttons) (move under menus?)
// Create Buttons (play, skip turn, rotation, etc)
// Main menu
// Pause state
// - LoadImageFromScreen, draw as tinted texture
// - Draw over that menus and the like
// Load Saves in window (Use s.GameState for saves)
// End game state and screen
// Pick starting player
// Skip Turns
//  -- auto-skip if player has no valid placements none and they skipped before
// Create random placement function to simulate games (from a given boardState)
// UI Work
// -- Button drawHover
// -- Ensure floating pieces are the right size (may simply need int cell Width/Height)
// -- (?) Center selected pieces on Cursor
// -- Flipping feedback (for when piece looks the same even with different 'orientation') (maybe do a slight gradient on the outline?)
// -- Abstract Drawing Boards (Game and Preview) into one function
// -- General Cleanup (Say which player's turn it is, better piece highlighting etc)
// -- (?) Enable line widths of > 1 (drawRect) (swapping over induces cell width errors, would need to address)
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

	// --------------------------------------------------
	// Main Game loop
	// --------------------------------------------------
	for !rl.WindowShouldClose() {
		game.StepGame()
	}
}
