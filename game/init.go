package game

import (
	s "github.com/AdamMcAdamson/blockeroni/state"
)

func Init() {
	s.Init()
	registerButtons()
}

func registerButtons() {
	registerMainMenuButtons()
	registerGameScreenButtons()
}

func registerMainMenuButtons() {
	//s.RegisterButton("MainMenuPlay")
}

func registerGameScreenButtons() {

}
