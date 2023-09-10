package gamestate

import (
	c "github.com/AdamMcAdamson/blockeroni/config"
)

// Get Pieces Structure Data
func getPiecesData() [21]c.PieceData {
	pieces := [21]c.PieceData{}

	// Cells
	// 1
	pieces[0].Cells = [][]bool{{true}}

	// 2
	pieces[1].Cells = [][]bool{{true, true}}

	// 3
	pieces[2].Cells = [][]bool{{false, true}, {true, true}}
	pieces[3].Cells = [][]bool{{true, true, true}}

	// 4
	pieces[4].Cells = [][]bool{{true, true, true, true}}
	pieces[5].Cells = [][]bool{{false, false, true}, {true, true, true}}
	pieces[6].Cells = [][]bool{{true, true, false}, {false, true, true}}
	pieces[7].Cells = [][]bool{{true, true}, {true, true}}
	pieces[8].Cells = [][]bool{{true, true, true}, {false, true, false}}

	// 5
	pieces[9].Cells = [][]bool{{false, true, true}, {true, true, false}, {false, true, false}}  // F
	pieces[10].Cells = [][]bool{{true}, {true}, {true}, {true}, {true}}                         // I
	pieces[11].Cells = [][]bool{{true, false}, {true, false}, {true, false}, {true, true}}      // L
	pieces[12].Cells = [][]bool{{false, true}, {true, true}, {true, false}, {true, false}}      // N
	pieces[13].Cells = [][]bool{{true, true}, {true, true}, {true, false}}                      // P
	pieces[14].Cells = [][]bool{{true, true, true}, {false, true, false}, {false, true, false}} // T
	pieces[15].Cells = [][]bool{{true, false, true}, {true, true, true}}                        // U
	pieces[16].Cells = [][]bool{{false, false, true}, {false, false, true}, {true, true, true}} // V
	pieces[17].Cells = [][]bool{{false, false, true}, {false, true, true}, {true, true, false}} // W
	pieces[18].Cells = [][]bool{{false, true, false}, {true, true, true}, {false, true, false}} // X
	pieces[19].Cells = [][]bool{{false, true}, {true, true}, {false, true}, {false, true}}      // Y
	pieces[20].Cells = [][]bool{{true, true, false}, {false, true, false}, {false, true, true}} // Z

	// Cells
	// 1
	pieces[0].Offset = [2]int{0, 0}

	// 2
	pieces[1].Offset = [2]int{0, 0}

	// 3
	pieces[2].Offset = [2]int{0, 0}
	pieces[3].Offset = [2]int{1, 0}

	// 4
	pieces[4].Offset = [2]int{1, 0}
	pieces[5].Offset = [2]int{1, 1}
	pieces[6].Offset = [2]int{1, 0}
	pieces[7].Offset = [2]int{0, 0}
	pieces[8].Offset = [2]int{1, 0}

	// 5
	pieces[9].Offset = [2]int{1, 1}  // F
	pieces[10].Offset = [2]int{0, 2} // I
	pieces[11].Offset = [2]int{0, 2} // L
	pieces[12].Offset = [2]int{0, 1} // N
	pieces[13].Offset = [2]int{0, 1} // P
	pieces[14].Offset = [2]int{1, 1} // T
	pieces[15].Offset = [2]int{1, 1} // U
	pieces[16].Offset = [2]int{2, 2} // V
	pieces[17].Offset = [2]int{1, 1} // W
	pieces[18].Offset = [2]int{1, 1} // X
	pieces[19].Offset = [2]int{1, 1} // Y
	pieces[20].Offset = [2]int{1, 1} // Z

	return pieces
}

var Pieces = getPiecesData()

func getPieceNumCellsFromIndex(index int) int {
	if index > 20 {
		panic("Invalid Index")
	} else if index >= 9 {
		return 5
	} else if index >= 4 {
		return 4
	} else if index >= 2 {
		return 3
	} else if index == 1 {
		return 2
	} else if index == 0 {
		return 1
	}
	panic("Invalid Index")
}
