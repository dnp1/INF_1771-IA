package templeFights

//import "fmt"
import "math"

type stateIndex struct {
	Previous *GameState
	Fighters int // fighters to get here.
}

//GameState ...
type GameState struct {
	Context   *context
	InxTemple int     // currentTemple
	TimeLeft  float64 // unsafe, but if works...
	Fighters  int     // fighters to get here.
	Previous  *GameState
	Lifes     []int
	Neighbors []*GameState
}

func (g *GameState) index() stateIndex {
	return stateIndex{g.Previous, g.Fighters}
}

// used to build neighbors
func (g *GameState) possibleFighters() []int {
	if g.Living() == 0 {
		return nil
	}
	return mapToPossibility[g.Living()]
}

func (g *GameState) Living() int {
	return encodeLiving(g.Lifes)
}

//CostToMe ...
func (g *GameState) CostToMe() float64 {
	var ctx = g.Context
	if g.InxTemple == 0 {
		return 0
	}
	difficulty := ctx.Temples[g.InxTemple-1].Difficulty
	if g.Fighters == 0 {
		return 0
	}
	var power = calcFightersPower(g.Fighters, ctx.Saints)
	return difficulty / power
}

//Quality ...
var optimalTime = float64(0)

func (g *GameState) Quality() float64 {
	var ctx = g.Context
	var difficulty = ctx.Temples[g.InxTemple-1].Difficulty
	//	var optimalTime = ctx.Temples
	var wantedPower = difficulty / optimalTime
	if g.Fighters == 0 {
		return 0
	}
	var power = calcFightersPower(g.Fighters, ctx.Saints)
	return (math.Abs(wantedPower-power) * optimalTime) / float64(g.InxTemple+1)

}

//EstimatedCost ...
func (g *GameState) EstimatedCost() float64 {
	var ctx = g.Context
	var difficultyLeft = float64(0)

	if g.Fighters == 0 {
		return 0
	}

	for inx := len(ctx.Temples) - 1 - g.InxTemple; inx >= 0; inx-- {
		difficultyLeft += ctx.Temples[inx].Difficulty
	}
	//fmt.Println("Dificuldade:", difficultyLeft)

	//	var wantedPower = difficulty / optimalTime

	var totalPower = float64(0) // calcFightersPower(g.Fighters, saints)
	for inx, saint := range ctx.Saints {
		totalPower += float64(g.Lifes[inx]) * saint.Power
	}

	//(math.Abs(wantedPower-power) * optimalTime * 50.0) / float64(g.InxTemple+1)

	//	fmt.Println("estimativa: ", difficultyLeft/totalPower)
	return difficultyLeft / totalPower // * float64((len(ctx.Temples) - g.InxTemple)) * g.Fighters
}

//IsFinal says if the state is a "acceptance" one
func (g *GameState) IsFinal() bool {
	return len(g.Context.Temples) == g.InxTemple && g.Living() != 0 && g.TimeLeft > 0
}
