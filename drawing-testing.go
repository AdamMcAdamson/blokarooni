package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() { // VS Code is very unhappy with this...
	rl.InitWindow(800, 450, "Raylib-go Drawing Test")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		drawSquareOnLaptop()

		// p1 := rl.Vector2{X: 10., Y: 10.}
		// p2 := rl.Vector2{X: 20., Y: 10.}
		// p3 := rl.Vector2{X: 20., Y: 20.}
		// p4 := rl.Vector2{X: 10., Y: 20.}

		// rl.DrawPixel(9, 9, rl.Green)
		// rl.DrawPixel(19, 9, rl.Green)
		// rl.DrawPixel(19, 19, rl.Green)
		// rl.DrawPixel(9, 19, rl.Green)

		/* // Remove Block comment to include grouped stuff

		rl.DrawLine(100, 50, 100, 200, rl.Green)
		rl.DrawLine(110, 50, 110, 200, rl.Green)

		// Draw Pixel Functions are Zero-Indexed

		rl.DrawPixel(100-1, 100-1, rl.Orange)
		rl.DrawPixel(110-1, 100-1, rl.Orange)
		rl.DrawPixel(100-1, 110-1, rl.Orange)
		rl.DrawPixel(110-1, 110-1, rl.Orange)

		// rl.DrawPixelV(rl.Vector2{X: 99.1, Y: 99.1}, rl.Green)
		// rl.DrawPixelV(rl.Vector2{X: 108.9, Y: 98.9}, rl.Green)
		// rl.DrawPixelV(rl.Vector2{X: 99.5, Y: 109.5}, rl.Green)  // Rounds up
		// rl.DrawPixelV(rl.Vector2{X: 108.5, Y: 108.5}, rl.Green) // Rounds Down

		// Draw Line Functions are not zero indexed (round down), and don't include the starting pixel (because they draw from a )

		rl.DrawLineV(rl.Vector2{X: 100.0, Y: 100.}, rl.Vector2{X: 110, Y: 100}, rl.Black)
		rl.DrawLine(100, 110, 110, 110, rl.Black)
		rl.DrawLine(110, 120, 100, 120, rl.Black)

		// rl.DrawLine(0, 0, 100, 0, rl.Black)
		// rl.DrawLine(0, 1, 100, 1, rl.Orange)
		// rl.DrawLine(1, 1, 100, 1, rl.Black)
		// rl.DrawLine(1, 1, 1, 1, rl.Black)

		// rl.DrawLineV(p1, p2, rl.Black)
		// rl.DrawLineV(p2, p3, rl.Red)
		// rl.DrawLineV(p3, p4, rl.Blue)
		// rl.DrawLineV(p4, p1, rl.Orange)

		// */

		rl.EndDrawing()
	}
}

func drawSquareOnLaptop() {

	// On my laptop:
	// Pixels are Zero-indexed

	// Orange (Lines)     :- Pixels: (0,0) - (0,49);  (0,49) - (49,49);  (0,0) - (49,0);  (49,0) - (49,49)
	// Red    (Pixels)    :- Pixels: (0,0);  (0,49);  (49,49); (49,0)
	// Black  (Rectangle) :- Pixels: (1,1) - (1,48);  (1,48) - (48,48);  (1,1) - (48,1);  (48,1) - (48,48); And everything in between

	rl.DrawLine(0, 1, 50, 1, rl.Orange)
	rl.DrawLine(1, 0, 1, 50, rl.Orange)
	rl.DrawLine(50, 0, 50, 50, rl.Orange)
	rl.DrawLine(0, 50, 50, 50, rl.Orange)

	rl.DrawPixel(0, 0, rl.Red)
	rl.DrawPixel(0, 49, rl.Red)
	rl.DrawPixel(49, 49, rl.Red)
	rl.DrawPixel(49, 0, rl.Red)

	rl.DrawRectangle(1, 1, 48, 48, rl.Black)
}
