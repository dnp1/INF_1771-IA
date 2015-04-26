package main

type GameState struct {
	inxTemple int     // currentTemple
	living    int     // encoded living saints
	timeLeft  float64 // unsafe, but if works...
}

// used to build neighbors
func (g *GameState) possibleFighters() [][]int {
	return mapToPossibility[g.living]
}

func (g *GameState) encode() {

}
