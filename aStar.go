package main

import (
	"fmt"
	"math"
)

// heuristicCostEstimate is, currently, just the distance between two pointers
func heuristicCostEstimate(origin *Square, goal *Square) int {
	var (
		dX       = float64(goal.Position.Row - origin.Position.Row)
		dY       = float64(goal.Position.Column - origin.Position.Column)
		distance = math.Sqrt(dX*dX + dY*dY)
	)

	if distance-math.Trunc(distance) > 0.5 {
		distance = distance + 1.0
	}

	return int(math.Floor(distance))
}

// getMin is a auxiliary function that return
func getMin(openSet map[*Square]bool, fScore map[*Square]int) *Square {
	var (
		best *Square
		min  int = 1<<30 - 1
	)

	for j, _ := range openSet {
		if fScore[j] <= min {
			min = fScore[j]
			best = j
		}
	}

	return best
}

// Return `reverse` Path.
func reconstructPath(cameFrom map[*Square]*Square, current *Square) ([]*Square, int) {
	var path = make([]*Square, 1, 42*42)
	var duration = int(0)
	path[0] = current
	duration = current.Cost()

	for next, ok := cameFrom[current]; ok && next != nil; next, ok = cameFrom[next] {
		path = append(path, next)
		duration = duration + next.Cost()
		current = next
	}

	// Fixing
	duration = duration - current.Cost()

	return path, duration
}

//AStar ...
func (v *Square) AStar(goal *Square) ([]*Square, int) {
	var (
		closedSet = make(map[*Square]bool)
		openSet   = map[*Square]bool{v: true}
		cameFrom  = make(map[*Square]*Square)
		gScore    = map[*Square]int{v: 0}
		fScore    = map[*Square]int{v: gScore[v] + heuristicCostEstimate(v, goal)}
	)
	fmt.Println("we are beginning!")
	for len(openSet) > 0 {

		var current = getMin(openSet, fScore)

		if current == goal {
			return reconstructPath(cameFrom, current)
		}

		delete(openSet, current)
		closedSet[current] = true

		for _, neighbor := range current.Neighbors() {
			if neighbor == nil || closedSet[neighbor] {
				continue
			}

			GScoreTry := gScore[current] + current.DistanceToNeighbor(neighbor)

			neighborInOpenSet := openSet[neighbor]

			if !neighborInOpenSet || GScoreTry < gScore[neighbor] {
				cameFrom[neighbor] = current
				gScore[neighbor] = GScoreTry
				fScore[neighbor] = gScore[neighbor] + heuristicCostEstimate(neighbor, goal)
				//adding it to openSet
				openSet[neighbor] = true
			}
		}

	}
	return nil, 0
}
