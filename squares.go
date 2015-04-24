package main

//import "errors"

// Square ..
type Square struct {
	Position  Point
	Cost      int64
	neighbors []*Square
}

//Neighbors .. .
func (v *Square) Neighbors() []*Square {
	return v.neighbors
}

//DistanceToNeighbor find a square in neighborhood and return their Cost. If "finding" is not a neighborhood, return -1
func (v *Square) DistanceToNeighbor(finding *Square) int64 {
	for _, neighbor := range v.Neighbors() {
		if neighbor == finding {
			return neighbor.Cost
		}
	}
	return -1
	//, errors.New("Not a Neighbor!")
}
