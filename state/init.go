package gamestate

import (
	c "github.com/AdamMcAdamson/blockeroni/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Init() {
	for i := range Players {
		Players[i].Id = i + 1
		Players[i].Turn = 0
		Players[i].PiecesRemaining = 21
		Players[i].Score = 0
		Players[i].Skipped = false
		Players[i].Pieces = [21]c.PieceState{}
		for j := range Players[i].Pieces {
			Players[i].Pieces[j].Number = j
			Players[i].Pieces[j].IsPlaced = false
			Players[i].Pieces[j].NumSquares = getPieceNumSquaresFromIndex(j)
		}
	}
	addPlayerSideboard(40, 120, 0)
	addPlayerSideboard(1360, 120, 1)
	addPlayerSideboard(40, 540, 2)
	addPlayerSideboard(1360, 540, 3)
}

func addPlayerSideboard(x int32, y int32, playerIndex int) {

	// margin of 5 in both dimensions
	drawRegionStartX := x + 20
	drawRegionStartY := y + 20

	posX := drawRegionStartX
	posY := drawRegionStartY

	// sideboardRec := rl.Rectangle{X: float32(x), Y: float32(y), Width: 400, Height: 380}
	// rl.DrawRectangleRounded(sideboardRec, 0.05, 1, rl.LightGray)

	for i, piece := range Players[playerIndex].Pieces {
		if !piece.IsPlaced {
			posX = drawRegionStartX + c.SideboardDrawOffsets[i][0]
			posY = drawRegionStartY + c.SideboardDrawOffsets[i][1]
			addPiece(posX, posY, 20, 20, i, playerIndex)
		}
	}

}

func addPiece(x int32, y int32, cellWidth int32, cellHeight int32, piece int, playerIndex int) {
	posX := x
	posY := y
	for py, prow := range Pieces[piece] {
		for px, pval := range prow {
			if pval {
				posX = x + (int32(px) * cellWidth)
				posY = y + (int32(py) * cellHeight)

				rect := rl.Rectangle{X: float32(posX), Y: float32(posY), Width: float32(cellWidth + 1), Height: float32(cellHeight + 1)}
				sideboardPieceSquare := c.SideboardPieceSquare{PieceNumber: piece, CollisionRect: rect}

				SideboardPieces[playerIndex] = append(SideboardPieces[playerIndex], sideboardPieceSquare)
			}
		}
	}
}
