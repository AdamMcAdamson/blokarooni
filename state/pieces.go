package gamestate

// Pieces Structure Data
func GetPiecesData() [21][][]bool {
	pieces := [21][][]bool{}

	pieces[0] = [][]bool{{true}}
	pieces[1] = [][]bool{{true, true}}
	pieces[2] = [][]bool{{false, true}, {true, true}}
	pieces[3] = [][]bool{{true, true, true}}
	pieces[4] = [][]bool{{true, true, true, true}}
	pieces[5] = [][]bool{{false, false, true}, {true, true, true}}
	pieces[6] = [][]bool{{true, true, false}, {false, true, true}}
	pieces[7] = [][]bool{{true, true}, {true, true}}
	pieces[8] = [][]bool{{true, true, true}, {false, true, false}}
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

var Pieces = GetPiecesData()
