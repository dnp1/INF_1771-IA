package templeFights

import (
	"fmt"
	//	e "github.com/daniloanp/IA/environment"
)

func backtrackedAStar(ctx *context, avTime float64) (bool, []*GameState) {
	var (
		states    = make(map[stateIndex]*GameState)
		closedSet = make(map[*GameState]bool)
		openSet   = make(map[*GameState]bool)
		cameFrom  = make(map[*GameState]*GameState)
		gScore    = make(map[*GameState]float64)
		fScore    = make(map[*GameState]float64)
		current   *GameState
		inxTemple = int(0)
		lifes     = make([]int, len(ctx.Saints))
	)

	for inx, saint := range ctx.Saints {
		lifes[inx] = saint.Lives
	}

	current = &GameState{
		Context:   ctx,
		InxTemple: inxTemple,
		TimeLeft:  avTime,
		Fighters:  0,
		Previous:  nil,
		Lifes:     lifes,
	}

	states[current.index()] = current

	openSet[current] = true
	gScore[current] = 0
	fScore[current] = gScore[current] /*+ /*current.EstimatedCost()*/

	fmt.Println("=======")

	for len(openSet) > 0 {
		var current = getMinState(openSet, fScore)

		if current.IsFinal() {
			return true, reconstructStatePath(cameFrom, current)
		}

		delete(openSet, current)
		closedSet[current] = true

		// visinhos
		for _, fighters := range current.possibleFighters() {
			var (
				difficulty = ctx.Temples[current.InxTemple].Difficulty
				cost       = calcCost(difficulty, fighters, ctx.Saints)
				neighbor   *GameState
			)
			index := stateIndex{current, fighters}
			if s := states[index]; s != nil {
				neighbor = s
			} else {
				neighbor = &GameState{
					Context:   ctx,
					InxTemple: current.InxTemple + 1,
					TimeLeft:  current.TimeLeft - cost,
					Fighters:  fighters,
					Previous:  current,
					Lifes:     make([]int, len(ctx.Saints)),
				}
				copy(neighbor.Lifes, current.Lifes)
				decreaseLifes(neighbor.Fighters, neighbor.Lifes)
				states[index] = neighbor
			}

			if closedSet[neighbor] {
				continue
			}

			GScoreTry := gScore[current] + cost

			if !openSet[neighbor] && neighbor.Living() != 0 && neighbor.TimeLeft > 0 || GScoreTry < gScore[neighbor] {
				cameFrom[neighbor] = current
				gScore[neighbor] = GScoreTry
				fScore[neighbor] = /*gScore[neighbor]*/ +neighbor.Quality()

				//adding it to openSet
				openSet[neighbor] = true
			}
		}

	}
	return false, nil
}

func getMinState(openSet map[*GameState]bool, fScore map[*GameState]float64) *GameState {
	var (
		best *GameState
		min  float64 = 1 << 31
	)

	for j, _ := range openSet {
		if fScore[j] <= min {
			min = fScore[j]
			best = j
		}
	}

	return best
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
