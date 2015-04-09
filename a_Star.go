package main

import "math"

//Currently is just the distance
func heuristicCostEstimate(origin *Square, goal *Square) int64 {
	var dX = float64(goal.p.X - origin.p.X)
	var dY = float64(goal.p.Y - origin.p.Y)

	var distance = math.Sqrt(dX*dX + dY*dY)

	if distance-math.Trunc(distance) > 0.5 {
		distance = distance + 1.0
	}
	return int64(math.Floor(distance))
}

// Return reverse Path.
func reconstructPath(cameFrom map[*Square]*Square, current *Square) []*Square {
	var path = make([]*Square, 1, 42*42)
	path[0] = current

	for next, ok := cameFrom[current]; ok; current = next {
		path = append(path, next)
	}
	return path
}

//AStar ...
func (v *Square) AStar(goal *Square) []*Square {
	var (
		closedSet = make(map[*Square]bool, 42*42)
		openSet   = map[*Square]bool{v: true}
		cameFrom  = make(map[*Square]*Square, 42*42)
		gScore    = map[*Square]int64{v: 0} // or v.Cost?
		fScore    = map[*Square]int64{v: gScore[v] + 0}
	)
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

			if !openSet[neighbor] || GScoreTry < gScore[neighbor] {
				cameFrom[neighbor] = current
				gScore[neighbor] = GScoreTry
				fScore[neighbor] = gScore[neighbor] + heuristicCostEstimate(neighbor, goal)
				//adding it to openSet
				openSet[neighbor] = true
			}
		}
	}
	return nil
}
