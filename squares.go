package main

import "fmt"

// Square ..
type Square struct {
	p         Point
	Cost      int64
	neighbors []*Square
}

//Neighbors .. .
func (v *Square) Neighbors() []*Square {
	return v.neighbors
}

func (v *Square) Equals(v1 *Square) bool {
	adj := v1
	return adj.p.X == v.p.X && adj.p.Y == v.p.Y
}

//CostToNeighbor ...
func (v *Square) CostToNeighbor(finding *Square) int64 {
	for _, adj := range v.Neighbors() {
		if adj == finding {
			return adj.Cost
		}
	}
	return -1
}

// BFS ...
func (v *Square) BFS(finding *Square) []*Square {
	// var cost = 0
	var Q = SquareQueue(make([]*Square, 0, 42*42))
	// var path = SquareQueue(make([]*Square, 0, 42*42))
	var D = make(map[*Square]bool)
	Q.add(v)
	D[v] = true

	for len(Q) > 0 {
		v := Q.get()
		if v == nil {
			break
		}
		for _, adj := range v.Neighbors() {
			if adj == nil {
				continue
			}
			if !D[adj] {
				if adj.Equals(finding) {
					return [](*Square){v, adj}
				}
				Q.add(adj)
				D[adj] = true
			}
		}
	}

	Q = nil
	return nil
}

func (v *Square) AStar(finding *Square) []*Square {
	var (
		closedSet = make(map[*Square]bool)
		openSet   = map[*Square]bool{v: true}
		cameFrom  = make(map[*Square]*Square)
		gScore    = map[*Square]int64{v: 0}
		fScore    = map[*Square]int64{v: (gScore[v] + v.CostToNeighbor(finding))}
	)

	for {

	}

	fmt.Println(closedSet, openSet, cameFrom, fScore)

	return nil
}

// function A*(start,goal)
//     closedset := the empty set    // The set of nodes already evaluated.
//     openset := {start}    // The set of tentative nodes to be evaluated, initially containing the start node
//     came_from := the empty map    // The map of navigated nodes.
//
//     g_score[start] := 0    // Cost from start along best known path.
//     // Estimated total cost from start to goal through y.
//     f_score[start] := g_score[start] + heuristic_cost_estimate(start, goal)
//
//     while openset is not empty
//         current := the node in openset having the lowest f_score[] value
//         if current = goal
//             return reconstruct_path(came_from, goal)
//
//         remove current from openset
//         add current to closedset
//         for each neighbor in neighbor_nodes(current)
//             if neighbor in closedset
//                 continue
//             tentative_g_score := g_score[current] + dist_between(current,neighbor)
//
//             if neighbor not in openset or tentative_g_score < g_score[neighbor]
//                 came_from[neighbor] := current
//                 g_score[neighbor] := tentative_g_score
//                 f_score[neighbor] := g_score[neighbor] + heuristic_cost_estimate(neighbor, goal)
//                 if neighbor not in openset
//                     add neighbor to openset
//
//     return failure
//
// function reconstruct_path(came_from,current)
//     total_path := [current]
//     while current in came_from:
//         current := came_from[current]
//         total_path.append(current)
//     return total_path
