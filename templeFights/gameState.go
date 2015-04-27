package templeFights

import "math"
import e "github.com/daniloanp/IA/environment"

type GameState struct {
	InxTemple int     // currentTemple
	Living    int     // encoded living saints
	TimeLeft  float64 // unsafe, but if works...
	Fighters  int     // fighters to get here.
	Previous  *GameState
}

// used to build neighbors
func (g *GameState) possibleFighters() []int {
	return mapToPossibility[g.Living]
}

func (g *GameState) CostToMe(saints []e.Saint) float64 {
	if g.InxTemple == 0 {
		return 0
	}
	difficulty := Temples[g.InxTemple-1].Difficulty
	if g.Fighters == 0 {
		return 0
	}
	var power = calcFightersPower(g.Fighters, saints)
	return truncFloat(difficulty / power)
}

func (g *GameState) Quality(saints []e.Saint) float64 {
	var difficulty = Temples[g.InxTemple-1].Difficulty
	var wantedPower = difficulty / optimalTime
	if g.Fighters == 0 {
		return 0
	}
	var power = calcFightersPower(g.Fighters, saints)
	return truncFloat((math.Abs(wantedPower-power) * optimalTime * 50.0) / float64(g.InxTemple+1))
}

func (g *GameState) IsFinal() bool {
	return len(Temples) == g.InxTemple && g.Living != 0
}
