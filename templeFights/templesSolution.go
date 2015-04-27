package templeFights

import (
	e "github.com/daniloanp/IA/environment"
	"math"
)

var optimalTime float64
var mapToPossibility MapToPossibility

var decodedFighters [][]int
var encodedFighters map[*[]int]int

var Temples []e.Temple

var Lives = make(map[*GameState][]int)

func TemplesSolution(saints []e.Saint, avTime float64, temples []e.Temple) (bool, []*GameState) {
	if mapToPossibility == nil {
		mapToPossibility.init(saints)
	}

	optimalTime = avTime / float64(len(temples)+1)

	Temples = temples

	return backtrackedAStar(saints, avTime)
}

func getMinState(openSet map[*GameState]bool, fScore map[*GameState]float64) *GameState {
	var (
		best *GameState
		min  float64 = 1<<30 - 1
	)

	for j, _ := range openSet {
		if fScore[j] <= min {
			min = fScore[j]
			best = j
		}
	}

	return best
}

func truncFloat(f float64) float64 {
	return math.Trunc(f*1000) / 1000

}
func reconstructStatePath(cameFrom map[*GameState]*GameState, current *GameState) []*GameState {
	var path = make([]*GameState, 1, 15)

	path[0] = current

	for next, ok := cameFrom[current]; ok; next, ok = cameFrom[next] {
		path = append(path, next)

		current = next
	}

	return path
}

func encodeLivingSaints(saints []e.Saint) int {
	var livingSaints = int(0)
	for inx, saint := range saints {
		if saint.Lives > 0 {
			livingSaints |= 1 << uint(inx)
		}
	}
	return livingSaints
}

func increaseSaints(fighters int, saints *[]e.Saint) {
	for inx, _ := range *saints {
		if (1<<uint(inx))&fighters > 0 {
			(*saints)[inx].Lives++
		}
	}
}

func decreaseLifes(fighters int, lifes []int) {
	for inx, _ := range lifes {
		if (1<<uint(inx))&fighters > 0 {
			lifes[inx]--
		}
	}
}

func encodeLiving(lifes []int) int {
	var living = int(0)

	for inx, lifes := range lifes {
		if lifes > 0 {
			living |= 1 << uint(inx)
		}
	}
	return living
}

var states = make(map[GameState]*GameState)

func getGameState(g GameState) *GameState {
	var v, ok = states[g]
	if ok && v != nil {
		return v
	}
	v = new(GameState)
	*v = g
	states[g] = v
	return v

}

func backtrackedAStar(saints []e.Saint, avTime float64) (bool, []*GameState) {
	var (
		closedSet = make(map[*GameState]bool)
		openSet   = make(map[*GameState]bool)
		cameFrom  = make(map[*GameState]*GameState)
		gScore    = make(map[*GameState]float64)
		fScore    = make(map[*GameState]float64)
		current   *GameState
		inxTemple = int(0)
	)

	var lives = make([]int, len(saints))
	for inx, saint := range saints {
		lives[inx] = saint.Lives
	}

	current = getGameState(GameState{
		Living:    encodeLivingSaints(saints),
		InxTemple: inxTemple,
		TimeLeft:  truncFloat(avTime),
		Fighters:  0,
		Previous:  nil,
	})

	openSet[current] = true
	gScore[current] = 0
	fScore[current] = gScore[current] + 0
	Lives[current] = lives

	for len(openSet) > 0 {
		var current = getMinState(openSet, fScore)

		if current.IsFinal() {
			return true, reconstructStatePath(cameFrom, current)

		}

		inxTemple = current.InxTemple
		avTime = current.TimeLeft

		delete(openSet, current)
		closedSet[current] = true
		var p = current.possibleFighters()

		var currentLifes = Lives[current]

		for _, fighters := range p {

			var difficulty = Temples[current.InxTemple].Difficulty
			var cost = calcCost(difficulty, fighters, saints)

			var lifes = make([]int, len(saints))

			copy(lifes, currentLifes)
			decreaseLifes(fighters, lifes)

			var neighbor = getGameState(GameState{
				Living:    encodeLiving(lifes),
				InxTemple: inxTemple + 1,
				TimeLeft:  truncFloat(avTime - cost),
				Fighters:  fighters,
				Previous:  current,
			})

			Lives[neighbor] = lifes

			if closedSet[neighbor] {
				continue
			}

			GScoreTry := gScore[current] + cost
			neighborInOpenSet := openSet[neighbor]

			if !neighborInOpenSet || GScoreTry < gScore[neighbor] {
				cameFrom[neighbor] = current
				gScore[neighbor] = GScoreTry
				fScore[neighbor] = /*gScore[neighbor] + */ neighbor.Quality(saints)

				//adding it to openSet
				openSet[neighbor] = true
			}
		}
	}
	return false, nil
}

func calcFightersPower(fighters int, saints []e.Saint) float64 {
	var sum = float64(0)

	for inx, saint := range saints {
		if (1<<uint(inx))&fighters > 0 {
			sum += saint.Power
		}
	}
	return sum
}

func calcCost(difficulty float64, fighters int, saints []e.Saint) float64 {
	var fightersPower = calcFightersPower(fighters, saints)
	return difficulty / fightersPower
}
