package main

// Vertex ...
type Vertex interface {
	Neighbors() []Vertex
	DistanceToNeighbor(v Vertex) int64
	BFS(finding Vertex) []Vertex
	Equals(v1 Vertex) bool
	// DFS(finding Vertex) []Vertex
}

// // Collection ...
// type Collection interface {
// 	Add(id uint64, val interface{}) bool
// 	Remove(id uint64) bool
// }
//
// func AddVet(c Collection, v []interface{}) {
// 	for i, val := range v {
// 		c.Add(uint64(i), val)
// 	}
// }
