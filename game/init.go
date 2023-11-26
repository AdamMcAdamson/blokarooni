package game

import (
	s "github.com/AdamMcAdamson/blockeroni/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Init() {
	s.Init()
	registerButtons()
	s.GameMode = 0
	setGameMode(0)
}

func registerButtons() {
	registerMainMenuButtons()
	registerGameScreenButtons()
}

func registerMainMenuButtons() {
	playRect := rl.Rectangle{X: 810, Y: 400, Width: 180, Height: 80}
	playOnPressed := func() {
		setGameMode(1)
	}
	playDrawUp := func(rect rl.Rectangle) {
		rl.DrawRectangleRounded(rect, 0.1, 4, rl.Gray)
		rl.DrawText("Play", rect.ToInt32().X+38, rect.ToInt32().Y+15, 48, rl.RayWhite)
	}
	playDrawDown := func(rect rl.Rectangle) {
		rl.DrawRectangleRounded(rect, 0.1, 4, rl.DarkGray)
		rl.DrawText("Play", rect.ToInt32().X+38, rect.ToInt32().Y+15, 48, rl.RayWhite)
	}
	s.RegisterButton("MainMenuPlay", playRect, playOnPressed, playDrawUp, playDrawDown)
}

func registerGameScreenButtons() {

}
