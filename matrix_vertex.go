package main

import "log"

//Matrix is a auxiliary type
//It is used to build a well connected "*Squares Graph" from a given *Environment
type Matrix struct {
	data         []*Square
	numOfRows    uint64
	numOfColumns uint64
}

//NewMatrix is a constructor for a *Matrix
func NewMatrix(numOfRows, numOfColumns uint64) *Matrix {
	return &Matrix{
		data:         make([]*Square, numOfRows*numOfColumns),
		numOfRows:    numOfRows,
		numOfColumns: numOfColumns,
	}
}

const (
	neighborTOP    = 0
	neighborRIGHT  = 1
	neighborBOTTOM = 2
	neighborLEFT   = 3
)

//note: the matrix is a Slice , then:
//get(i,0)-> i*mat.m + 0 -> walks i times the m coluns-in the slice-reaching the (i,j) element
//get(i,j) -> i*mat.m + j -> walks i times the m coluns-in the slice- and walks j elements, to reach(i,j)element
func (mat *Matrix) get(row, column uint64) *Square {
	if row*mat.numOfRows+column < uint64(len(mat.data)) && column < mat.numOfColumns {
		return mat.data[row*mat.numOfRows+column]
	}
	return nil
}

//The set method defines a value do a given position (as a [row, column] coordinate)
func (mat *Matrix) set(row uint64, column uint64, vertexInfo *Square) {
	mat.data[row*mat.numOfRows+column] = vertexInfo
}

// getCost find the fixed-cost to move to a square.
func getCost(env *Environment, squareID string, x uint64, y uint64) int64 {
	switch {
	case env.Start.X == x && env.Start.Y == y: // is start
		fallthrough // The same action below
	case env.End.X == x && env.End.Y == y: // is the stop
		return 0
	case squareID == "_": // Must be a temple square. otherwise, and error will be emitted!
		for _, t := range env.Temples { //return cost of the current temple
			if t.Position.X == uint64(x) && t.Position.Y == uint64(y) {
				return t.Cost
			}
		}
	default: // Must be a normal ground square, otherwise, and error will be emitted!
		for _, g := range env.Grounds { //given a square, searches in env the ground returning the cost of c
			//if t.Position.X == uint64(x) && t.Position.Y == uint64(y) {
			if g.ID == squareID {
				return g.Cost
			}
		}
		log.Fatalln("Invalid Ground at", x, y, "!!!\n")
	}
	return 0
}

// getOrBuild function return the same as "Matrix.get" but:
// it building the info case it was nil
// it need of env to build env
func getOrBuild(env *Environment, ref *Matrix, x uint64, y uint64) *Square {
	if v := ref.get(x, y); v != nil { // If v already are defined we just return it.
		return v
	} // Otherwise, we need build it.
	s := new(Square)
	s.Cost = getCost(env, env.Map[x][y], x, y)
	s.Position = Point{X: x, Y: y}
	s.neighbors = make([]*Square, 4)
	ref.set(x, y, s)
	return s
}

// buildGraphFromEnv build a Graph and return the initial and a slice with goals-squares
func buildGraphFromEnv(env *Environment) (*Square, []*Square) {
	var (
		ref          = NewMatrix(42, 42)
		destinations = make([]*Square, 0, len(env.Temples)+1)
	)

	for x, l := range env.Map {
		for y, _ := range l {
			var (
				x, y = uint64(x), uint64(y)
				s    = getOrBuild(env, ref, x, y)
			)

			if x > 0 { // has top neighbor
				s.neighbors[neighborTOP] = getOrBuild(env, ref, x-1, y) // this shouldn't be... x, y-1) instead ...x-1,y)? after all, you will the neighbor on top, this mean y -1 by convention
			}
			if y > 0 { // has  left neighbor
				s.neighbors[neighborLEFT] = getOrBuild(env, ref, x, y-1)
			}
			if x < uint64(len(env.Map)-1) { // has bottom neighbor
				s.neighbors[neighborBOTTOM] = getOrBuild(env, ref, x+1, y)
			}
			if y < uint64(len(l)-1) { // has right neighbor
				s.neighbors[neighborRIGHT] = getOrBuild(env, ref, x, y+1)
			}
		}
	}

	//filling "destinations" assuming env.Temples are in-order of goalss
	for _, v := range env.Temples {
		destinations = append(destinations, ref.get(v.Position.X, v.Position.Y))
	}
	destinations = append(destinations, ref.get(env.End.X, env.End.Y))

	return ref.get(env.Start.X, env.Start.Y), destinations
}
