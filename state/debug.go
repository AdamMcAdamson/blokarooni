package gamestate

import (
	"fmt"
)

func DebugPrint() {
	fmt.Printf("%v\n", BoardState)
	fmt.Printf("CurrentPlayerIndex: %d, PieceToPlace: %d\n", CurrentPlayerIndex, PieceToPlace)
}
