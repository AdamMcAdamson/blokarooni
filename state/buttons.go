package gamestate

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type fn func()
type drawfn func(rl.Rectangle)

type button struct {
	name        string
	active      bool
	clickedDown bool
	rect        rl.Rectangle
	onPressed   fn
	drawUp      drawfn
	drawDown    drawfn
}

var buttons []button

func RegisterButton(name string, rect rl.Rectangle, onPressed fn, drawUp drawfn, drawDown drawfn) bool {
	for i := range buttons {
		if buttons[i].name == name {
			return false
		}
	}
	buttons = append(buttons, button{name: name, active: false, clickedDown: false, rect: rect, onPressed: onPressed, drawUp: drawUp, drawDown: drawDown})
	return true
}

func DisableButton(name string) (bool, string) {
	for i := range buttons {
		if buttons[i].name == name {
			buttons[i].active = false
			return true, "success"
		}
	}
	return false, fmt.Sprintf("Button %s not found", name)
}

func DisableAllButtons() {
	for i := range buttons {
		buttons[i].active = false
		buttons[i].clickedDown = false
	}
}

// func releaseAllButtons() {
// 	for i := range buttons {
// 		buttons[i].clickedDown = false
// 	}
// }

func EnableButton(name string) (bool, string) {
	var buttonIndex int
	for i := range buttons {
		if buttons[i].name == name {
			buttonIndex = i
			break
		}
	}
	for i := range buttons {
		if i != buttonIndex && rl.CheckCollisionRecs(buttons[i].rect, buttons[buttonIndex].rect) {
			return false, fmt.Sprintf("Conflict with existing button %s", buttons[i].name)
		}
	}
	buttons[buttonIndex].active = true
	return true, "success"
}

func DetectAndHandleButtonDown(mousePos rl.Vector2) (bool, string) {
	for i := range buttons {
		if buttons[i].active && rl.CheckCollisionPointRec(mousePos, buttons[i].rect) {
			buttons[i].clickedDown = true
			return true, buttons[i].name
		}
	}
	return false, ""
}

func DetectAndHandleButtonRelease(mousePos rl.Vector2) (bool, string) {
	for i := range buttons {
		if buttons[i].active && buttons[i].clickedDown && rl.CheckCollisionPointRec(mousePos, buttons[i].rect) {
			buttons[i].onPressed()
			buttons[i].clickedDown = false
			return true, buttons[i].name
		}
		buttons[i].clickedDown = false
	}
	return false, ""
}

func DrawActiveButtons(mousePos rl.Vector2) {
	for i := range buttons {
		if buttons[i].active {
			if buttons[i].clickedDown && rl.CheckCollisionPointRec(mousePos, buttons[i].rect) {
				buttons[i].drawDown(buttons[i].rect)
			} else {
				buttons[i].drawUp(buttons[i].rect)
			}
		}
	}
}
