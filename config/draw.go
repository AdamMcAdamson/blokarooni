package config

func GenerateSideboardDrawOffsets(cellWidth int32, cellHeight int32 /*, sideboardWidth int32, sideboardHeight int32*/) [21][2]int32 {
	workingOffsets := [21][2]int32{
		// row 1
		{0, 0},
		{cellWidth * 2, 0},
		{cellWidth * 5, 0},
		{cellWidth * 8, 0},
		{cellWidth * 12, 0},

		// row 2
		{0, cellHeight * 3},
		{cellWidth * 4, cellHeight * 3},
		{cellWidth * 8, cellHeight * 3},
		{cellWidth * 11, cellHeight * 3},
		{cellWidth * 15, cellHeight * 3},

		// row 3
		{0, cellHeight * 7},
		{cellWidth * 2, cellHeight * 7},
		{cellWidth * 5, cellHeight * 7},
		{cellWidth * 8, cellHeight * 7},
		{cellWidth * 11, cellHeight * 7},
		{cellWidth * 15, cellHeight * 7},

		// row 4
		{0, cellHeight * 13},
		{cellWidth * 4, cellHeight * 13},
		{cellWidth * 8, cellHeight * 13},
		{cellWidth * 12, cellHeight * 13},
		{cellWidth * 15, cellHeight * 13},
	}

	return workingOffsets
}

var SideboardDrawOffsets = GenerateSideboardDrawOffsets(20, 20)
