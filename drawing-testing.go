package main

/*
import rl "github.com/gen2brain/raylib-go/raylib"

func main() { // VS Code is very unhappy with this...
	rl.InitWindow(450, 450, "Raylib-go Drawing Test")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		// drawOnLaptop()
		drawOnDesktop()

		// p1 := rl.Vector2{X: 10., Y: 10.}
		// p2 := rl.Vector2{X: 20., Y: 10.}
		// p3 := rl.Vector2{X: 20., Y: 20.}
		// p4 := rl.Vector2{X: 10., Y: 20.}

		// rl.DrawPixel(9, 9, rl.Green)
		// rl.DrawPixel(19, 9, rl.Green)
		// rl.DrawPixel(19, 19, rl.Green)
		// rl.DrawPixel(9, 19, rl.Green)

		/*
		// Remove Block comment to include grouped stuff

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

		// */ /*


		rl.EndDrawing()
	}
}

func drawOnLaptop() {
	drawSquareOnLaptop()
	drawPixelsAndLinesOnLaptop()
}

func drawOnDesktop() {
	drawLinesExample(200, 10, 1)
	drawSquareOnDesktop()
	drawPixelsOnDesktop()
	drawLinesOnDesktop()
}

func drawLinesExample(offsetX int32, offsetY int32, spacingX int32) {
	ex := offsetX
	ey := offsetY

	if spacingX < 6 {
		spacingX = 6
	}

	// 45 Deg Diagonal: behaves as expected
	// ------------------------
	rl.DrawPixel(0+ex, 0+ey, rl.Black)
	rl.DrawPixel(1+ex, 1+ey, rl.Black)
	rl.DrawPixel(2+ex, 2+ey, rl.Black)
	rl.DrawPixel(3+ex, 3+ey, rl.Black)
	rl.DrawPixel(4+ex, 4+ey, rl.Black)

	rl.DrawLine(1+ex, 1+ey, 4+ex, 4+ey, rl.Red)
	// ------------------------

	ex += spacingX

	// Horizontal (Exacting): behaves as expected
	// ------------------------
	rl.DrawPixel(0+ex, 1+ey, rl.Black)
	rl.DrawPixel(1+ex, 1+ey, rl.Black)
	rl.DrawPixel(2+ex, 1+ey, rl.Black)
	rl.DrawPixel(3+ex, 1+ey, rl.Black)
	rl.DrawPixel(4+ex, 1+ey, rl.Black)

	rl.DrawLine(1+ex, 1+ey, 4+ex, 2+ey, rl.Red)
	// ------------------------

	ex += spacingX

	// Horizontal (Ambuguous): Y draws higher pixel value, filling the pixel below the line index
	// ------------------------
	rl.DrawPixel(0+ex, 1+ey, rl.Black)
	rl.DrawPixel(1+ex, 1+ey, rl.Black)
	rl.DrawPixel(2+ex, 1+ey, rl.Black)
	rl.DrawPixel(3+ex, 1+ey, rl.Black)
	rl.DrawPixel(4+ex, 1+ey, rl.Black)

	rl.DrawLine(1+ex, 1+ey, 4+ex, 1+ey, rl.Red) // Y draws higher pixel value, filling the pixel below the line index
	// ------------------------

	ex += spacingX

	// Vertical (Exacting): behaves as expected
	// ------------------------
	rl.DrawPixel(1+ex, 0+ey, rl.Black)
	rl.DrawPixel(1+ex, 1+ey, rl.Black)
	rl.DrawPixel(1+ex, 2+ey, rl.Black)
	rl.DrawPixel(1+ex, 3+ey, rl.Black)
	rl.DrawPixel(1+ex, 4+ey, rl.Black)

	rl.DrawLine(1+ex, 1+ey, 2+ex, 4+ey, rl.Red)
	// ------------------------

	ex += spacingX

	// Vertical (ambuguous): X draws the lower pixel value, filling the pixel left of the line index
	// ------------------------
	rl.DrawPixel(1+ex, 0+ey, rl.Black)
	rl.DrawPixel(1+ex, 1+ey, rl.Black)
	rl.DrawPixel(1+ex, 2+ey, rl.Black)
	rl.DrawPixel(1+ex, 3+ey, rl.Black)
	rl.DrawPixel(1+ex, 4+ey, rl.Black)

	rl.DrawLine(1+ex, 1+ey, 1+ex, 4+ey, rl.Red) // X draws the lower pixel value, filling the pixel left of the line index
	// ------------------------
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

func drawSquareOnDesktop() {

	// On my laptop:
	// Pixels are Zero-indexed

	// Orange (Lines)     :- Pixels: (0,0) - (0,49);  (0,49) - (49,49);  (0,0) - (49,0);  (49,0) - (49,49)
	// Red    (Pixels)    :- Pixels: (0,0);  (0,49);  (49,49); (49,0)
	// Black  (Rectangle) :- Pixels: (1,1) - (1,48);  (1,48) - (48,48);  (1,1) - (48,1);  (48,1) - (48,48); And everything in between

	// rl.DrawPixel(0, 0, rl.Red)

	rl.DrawLine(0, 0, 50, 0, rl.Orange)   // Pixels: ( 0, 0) - (49, 0)
	rl.DrawLine(1, 0, 1, 50, rl.Orange)   // Pixels: ( 0, 0) - ( 0,49)
	rl.DrawLine(50, 0, 50, 50, rl.Orange) // Pixels: (49, 0) - (49,49)
	rl.DrawLine(0, 49, 50, 49, rl.Orange) // Pixels: ( 0,49) - (49,49)

	rl.DrawPixel(0, 0, rl.Red)
	rl.DrawPixel(0, 49, rl.Red)
	rl.DrawPixel(49, 49, rl.Red)
	rl.DrawPixel(49, 0, rl.Red)

	rl.DrawRectangle(1, 1, 48, 48, rl.Black)
}

func drawPixelsAndLinesOnLaptop() {

	// p1 := rl.Vector2{X: 10., Y: 10.}
	// p2 := rl.Vector2{X: 20., Y: 10.}
	// p3 := rl.Vector2{X: 20., Y: 20.}
	// p4 := rl.Vector2{X: 10., Y: 20.}

	// rl.DrawPixel(9, 9, rl.Green)
	// rl.DrawPixel(19, 9, rl.Green)
	// rl.DrawPixel(19, 19, rl.Green)
	// rl.DrawPixel(9, 19, rl.Green)

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

	// */ /*

}

func drawPixelsOnDesktop() {
	// Vector rounding behaves differently for X and Y: X Rounds Down, Y Rounds Up

	rl.DrawPixel(200, 200, rl.Black)
	rl.DrawPixel(210, 200, rl.Black)
	rl.DrawPixel(200, 210, rl.Black)
	rl.DrawPixel(210, 210, rl.Black)

	rl.DrawPixelV(rl.Vector2{X: 200.4, Y: 200.4}, rl.Green) // (200, 200)
	rl.DrawPixelV(rl.Vector2{X: 209.6, Y: 199.6}, rl.Green) // (210, 200)
	rl.DrawPixelV(rl.Vector2{X: 200.5, Y: 210.5}, rl.Green) // (200, 211) :- X Rounds Down, Y Rounds Up
	rl.DrawPixelV(rl.Vector2{X: 209.5, Y: 209.5}, rl.Green) // (209, 211) :- X Rounds Down, Y Rounds Up (even with the same float value!! WTF is that?!)

}

func drawLinesOnDesktop() {

	// Here, X behaves as zero-indexed, Y behaves as one-indexed

	rl.DrawLine(100, 50, 100, 200, rl.Green) // Pixels: ( 99, 50) - ( 99, 199)
	rl.DrawLine(110, 200, 110, 50, rl.Green) // Pixels: (109, 50) - (109, 199)

	// Draw Pixel Functions are Zero-Indexed

	rl.DrawPixel(99, 50, rl.Orange)   // First Pixel 			(Line 1)
	rl.DrawPixel(109, 51, rl.Orange)  // Second Pixel 			(Line 2)
	rl.DrawPixel(99, 199, rl.Orange)  // Last Pixel				(Line 1)
	rl.DrawPixel(109, 198, rl.Orange) // Second to Last Pixel	(Line 2)

	rl.DrawPixel(100-1, 100, rl.Orange) // ( 99, 100)
	rl.DrawPixel(110-1, 100, rl.Orange) // (109, 100)
	rl.DrawPixel(100-1, 110, rl.Orange) // ( 99, 110)
	rl.DrawPixel(110-1, 110, rl.Orange) // (109, 110)
	rl.DrawPixel(100-1, 120, rl.Orange) // ( 99, 120)
	rl.DrawPixel(110-1, 120, rl.Orange) // (109, 120)

	rl.DrawLineV(rl.Vector2{X: 100., Y: 100.}, rl.Vector2{X: 110, Y: 100}, rl.Black) // Pixels: (100,100) - (109, 100)
	rl.DrawLine(100, 110, 110, 110, rl.Black)                                        // Pixels: (100,110) - (109, 110)
	rl.DrawLine(110, 120, 100, 120, rl.Black)                                        // Pixels: (100,120) - (109, 120)

}
*/
