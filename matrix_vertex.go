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
func getCost(env *Environment, squareID string, row uint64, column uint64) int64 {
	switch {
	case env.Start.Row == row && env.Start.Column == column: // is start
		fallthrough // The same action below
	case env.End.Row == row && env.End.Column == column: // is the stop
		return 0
	case squareID == "_": // Must be a temple square. otherwise, and error will be emitted!
		for _, t := range env.Temples { //return cost of the current temple
			if t.Position.Row == uint64(row) && t.Position.Column == uint64(column) {
				return t.Cost
			}
		}
		log.Fatalln("Invalid Temple at", row, column, "!!!\n")
	default: // Must be a normal ground square, otherwise, and error will be emitted!
		for _, g := range env.Grounds { //given a square, searches in env the ground returning the cost of c
			if g.ID == squareID {
				return g.Cost
			}
		}
		log.Fatalln("Invalid Ground at", row, column, "!!!\n")
	}
	return 0
}

// getOrBuild function return the same as "Matrix.get" but:
// it building the info case it was nil
// it need of env to build env
func getOrBuild(env *Environment, ref *Matrix, row uint64, column uint64) *Square {
	if v := ref.get(row, column); v != nil { // If v already are defined we just return it.
		return v
	} // Otherwise, we need build it.
	s := new(Square)
	s.Cost = getCost(env, env.Map[row][column], row, column)
	s.Position = Point{Row: row, Column: column}
	s.neighbors = make([]*Square, 4)
	ref.set(row, column, s)
	return s
}

// buildGraphFromEnv build a Graph and return the initial and a slice with goals-squares
func buildGraphFromEnv(env *Environment) (*Square, []*Square, *Matrix) {
	var (
		ref          = NewMatrix(42, 42)
		destinations = make([]*Square, 0, len(env.Temples)+1)
	)

	for row, l := range env.Map {
		for column, _ := range l {
			var (
				row, column   = uint64(row), uint64(column)
				currentSquare = getOrBuild(env, ref, row, column)
			)

			if row > 0 { // has top neighbor
				currentSquare.neighbors[neighborTOP] = getOrBuild(env, ref, row-1, column)
			}
			if column > 0 { // has  left neighbor
				currentSquare.neighbors[neighborLEFT] = getOrBuild(env, ref, row, column-1)
			}
			if row < uint64(len(env.Map)-1) { // has bottom neighbor
				currentSquare.neighbors[neighborBOTTOM] = getOrBuild(env, ref, row+1, column)
			}
			if column < uint64(len(l)-1) { // has right neighbor
				currentSquare.neighbors[neighborRIGHT] = getOrBuild(env, ref, row, column+1)
			}
		}
	}

	//filling "destinations" assuming env.Temples are in-order of goalss
	for _, v := range env.Temples {
		destinations = append(destinations, ref.get(v.Position.Row, v.Position.Column))
	}
	destinations = append(destinations, ref.get(env.End.Row, env.End.Column))

	return ref.get(env.Start.Row, env.Start.Column), destinations, ref
}
