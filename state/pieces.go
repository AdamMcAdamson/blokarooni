package gamestate

// Pieces Structure Data
func getPiecesData() [21][][]bool {
	pieces := [21][][]bool{}

	// 1
	pieces[0] = [][]bool{{true}}

	// 2
	pieces[1] = [][]bool{{true, true}}

	// 3
	pieces[2] = [][]bool{{false, true}, {true, true}}
	pieces[3] = [][]bool{{true, true, true}}

	// 4
	pieces[4] = [][]bool{{true, true, true, true}}
	pieces[5] = [][]bool{{false, false, true}, {true, true, true}}
	pieces[6] = [][]bool{{true, true, false}, {false, true, true}}
	pieces[7] = [][]bool{{true, true}, {true, true}}
	pieces[8] = [][]bool{{true, true, true}, {false, true, false}}

	// 5
	pieces[9] = [][]bool{{false, true, true}, {true, true, false}, {false, true, false}}  // F
	pieces[10] = [][]bool{{true}, {true}, {true}, {true}, {true}}                         // I
	pieces[11] = [][]bool{{true, false}, {true, false}, {true, false}, {true, true}}      // L
	pieces[12] = [][]bool{{false, true}, {true, true}, {true, false}, {true, false}}      // N
	pieces[13] = [][]bool{{true, true}, {true, true}, {true, false}}                      // P
	pieces[14] = [][]bool{{true, true, true}, {false, true, false}, {false, true, false}} // T
	pieces[15] = [][]bool{{true, false, true}, {true, true, true}}                        // U
	pieces[16] = [][]bool{{false, false, true}, {false, false, true}, {true, true, true}} // V
	pieces[17] = [][]bool{{false, false, true}, {false, true, true}, {true, true, false}} // W
	pieces[18] = [][]bool{{false, true, false}, {true, true, true}, {false, true, false}} // X
	pieces[19] = [][]bool{{false, true}, {true, true}, {false, true}, {false, true}}      // Y
	pieces[20] = [][]bool{{true, true, false}, {false, true, false}, {false, true, true}} // Z

	return pieces
}

var Pieces = getPiecesData()

func getPieceNumSquaresFromIndex(index int) int {
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
