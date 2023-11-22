package gamestate

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type fn func()
type drawfn func(rl.Rectangle)

type button struct {
	name      string
	active    bool
	rect      rl.Rectangle
	onPressed fn
	draw      drawfn
}

var buttons []button

func RegisterButton(name string, rect rl.Rectangle, onPressed fn, draw drawfn) bool {
	for i := range buttons {
		if buttons[i].name == name {
			return false
		}
	}
	buttons = append(buttons, button{name: name, active: false, rect: rect, onPressed: onPressed, draw: draw})
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
	}
}

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

func DetectAndHandleButtonPress(location rl.Vector2) (bool, string) {
	for i := range buttons {
		if buttons[i].active && rl.CheckCollisionPointRec(location, buttons[i].rect) {
			buttons[i].onPressed()
			return true, buttons[i].name
		}
	}
	return false, ""
}

func DrawActiveButtons() {
	for i := range buttons {
		if buttons[i].active {
			buttons[i].draw(buttons[i].rect)
		}
	}
}
