package main

import (
	"fmt"
	"math"
)

var currentDistance float64

//The math definition of distance
func distanceBetweenTwoPoints(p1, p2 Point) float64 {
	var (
		dX       = float64(p1.Row - p2.Row)
		dY       = float64(p1.Column - p2.Column)
		distance = math.Sqrt(dX*dX + dY*dY)
	)
	return distance
}

// heuristicCostEstimate for AStar
func heuristicCostEstimate(origin *Square, goal *Square) float64 {
	var distance = distanceBetweenTwoPoints(origin.Position, goal.Position)

	if distance >= currentDistance {
		var min int = 1 << 20
		for _, neighbor := range origin.Neighbors() { // which is the neighbor nearest to goal with min cost
			if distanceBetweenTwoPoints(neighbor.Position, goal.Position) < distance && neighbor.Cost() < min {
				min = neighbor.Cost()
			}
		}
		distance += float64(min)
	}
	return distance
}

// getMin is a auxiliary function that return
func getMin(openSet map[*Square]bool, fScore map[*Square]float64) *Square {
	var (
		best *Square
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

// Return `reverse` Path.
func reconstructPath(cameFrom map[*Square]*Square, current *Square) []*Square {
	var path = make([]*Square, 1, 42*42)

	path[0] = current

	for next, ok := cameFrom[current]; ok && next != nil; next, ok = cameFrom[next] {
		path = append(path, next)

		current = next
	}

	return path
}

//AStar ...
func (v *Square) AStar(goal *Square) ([]*Square, int) {
	var (
		closedSet = make(map[*Square]bool)
		openSet   = map[*Square]bool{v: true}
		cameFrom  = make(map[*Square]*Square)
		gScore    = map[*Square]float64{v: 0}
		fScore    = map[*Square]float64{v: gScore[v] + heuristicCostEstimate(v, goal)}
	)
	fmt.Println("we are beginning!")
	for len(openSet) > 0 {

		var current = getMin(openSet, fScore)

		if current == goal {
			return reconstructPath(cameFrom, current), int(gScore[current])
		}

		delete(openSet, current)
		closedSet[current] = true

		currentDistance = distanceBetweenTwoPoints(current.Position, goal.Position)

		for _, neighbor := range current.Neighbors() {
			if neighbor == nil || closedSet[neighbor] {
				continue
			}

			GScoreTry := gScore[current] + float64(neighbor.Cost())

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
