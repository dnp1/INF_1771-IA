package pathThroughMap

import (
	e "github.com/daniloanp/IA/environment"
	"log"
)

type TempleInfo struct {
	Name       string
	Difficulty float64
}
type GroundInfo struct {
	ID   string
	Cost int
}
type Square struct {
	Position   e.Point
	TempleData *TempleInfo
	GroundData *GroundInfo
	neighbors  []*Square
}

//Neighbors .. .
func (v *Square) Neighbors() []*Square {
	return v.neighbors
}

func (v *Square) Cost() int {
	if v.GroundData != nil {
		return v.GroundData.Cost
	} else if v.TempleData != nil {
		return 0
	}
	return 1
}

func (v *Square) IsTemple() bool {
	return v.TempleData != nil
}

//DistanceToNeighbor find a square in neighborhood and return their Cost. If "finding" is not a neighborhood, return -1
func (v *Square) DistanceToNeighbor(finding *Square) int {
	for _, neighbor := range v.Neighbors() {
		if neighbor == finding {
			return neighbor.Cost()
		}
	}
	log.Fatalln("not a neighbor")
	return -1
}
