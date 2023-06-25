package gamestate

func Init() {
	for i := range Players {
		for j := range Players[i].Pieces {
			Players[i].Pieces[j].Number = j
			Players[i].Pieces[j].IsPlaced = false
		}
	}
}
