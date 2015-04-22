package main

// Vertex is a representation of a ""
type Vertex interface {
	Neighbors() []Vertex
	DistanceToNeighbor(v Vertex) int64
	BFS(finding Vertex) []Vertex
	Equals(v1 Vertex) bool
	// DFS(finding Vertex) []Vertex
}
