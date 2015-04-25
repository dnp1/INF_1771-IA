package main

import "log"

//import "fmt"

//Matrix is a auxiliary type
//It is used to build a well connected "*Squares Graph" from a given *Environment
type Matrix struct {
	data         []*Square
	numOfRows    int
	numOfColumns int
}

//NewMatrix is a constructor for a *Matrix
func NewMatrix(numOfRows, numOfColumns int) *Matrix {
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
func (mat *Matrix) get(row, column int) *Square {
	if row*mat.numOfRows+column < int(len(mat.data)) && column < mat.numOfColumns {
		return mat.data[row*mat.numOfRows+column]
	}
	return nil
}

//The set method defines a value do a given position (as a [row, column] coordinate)
func (mat *Matrix) set(row int, column int, vertexInfo *Square) {
	mat.data[row*mat.numOfRows+column] = vertexInfo
}

func getTempleData(env *Environment, squareID string, row int, column int) *TempleInfo {
	if squareID != "_" {
		return nil
	} // otherwise, must be a  valid Temple
	for _, t := range env.Temples { //return cost of the current temple
		if t.Position.Row == int(row) && t.Position.Column == int(column) {
			return &TempleInfo{
				Name:       t.Name,
				Difficulty: t.Difficulty,
			}
		}
	}
	log.Fatalln("Invalid Temple at", row, column, "!!!\n")
	return nil
}

func getGroundData(env *Environment, squareID string, row int, column int) *GroundInfo {
	if squareID == "_" {

		return nil
	}
	for _, g := range env.Grounds { //given a square, searches in env the ground returning the cost of c
		if g.ID == squareID {
			return &GroundInfo{
				ID:   g.ID,
				Cost: g.Cost,
			}
		}
	}
	log.Fatalln("Invalid Ground at", row, column, "!!!\n", squareID)
	return nil
}

// getOrBuild function return the same as "Matrix.get" but:
// it building the info case it was nil
// it need of env to build env
func getOrBuild(env *Environment, ref *Matrix, row int, column int) *Square {
	if v := ref.get(row, column); v != nil { // If v already are defined we just return it.
		return v
	} // Otherwise, we need build it.
	s := new(Square)
	// s.GroundData.Cost = getCost(env, env.Map[row][column], row, column)
	if !(env.End.Row == row && env.End.Column == column || env.Start.Row == row && env.Start.Column == column) {
		s.GroundData = getGroundData(env, env.Map[row][column], row, column)
		s.TempleData = getTempleData(env, env.Map[row][column], row, column)
	}
	s.Position = Point{Row: row, Column: column}
	s.neighbors = make([]*Square, 4)
	ref.set(row, column, s)
	return s
}

//buildGraphFromEnvironment build a Graph and return the initial and a slice with goals-squares
func buildGraphFromEnvironment(env *Environment) (*Square, []*Square, *Matrix) {
	var (
		ref          = NewMatrix(42, 42)
		destinations = make([]*Square, 0, len(env.Temples)+1)
	)

	for row, l := range env.Map {
		for column, _ := range l {
			var (
				row, column   = int(row), int(column)
				currentSquare = getOrBuild(env, ref, row, column)
			)

			if row > 0 { // has top neighbor
				currentSquare.neighbors[neighborTOP] = getOrBuild(env, ref, row-1, column)
			}
			if column > 0 { // has  left neighbor
				currentSquare.neighbors[neighborLEFT] = getOrBuild(env, ref, row, column-1)
			}
			if row < int(len(env.Map)-1) { // has bottom neighbor
				currentSquare.neighbors[neighborBOTTOM] = getOrBuild(env, ref, row+1, column)
			}
			if column < int(len(l)-1) { // has right neighbor
				currentSquare.neighbors[neighborRIGHT] = getOrBuild(env, ref, row, column+1)
			}
		}
	}

	//filling "destinations" assuming env.Temples are in-order of goalss
	for _, v := range env.Temples {
		destinations = append(destinations, ref.get(v.Position.Row, v.Position.Column))
	}
	destinations = append(destinations, ref.get(env.End.Row, env.End.Column))

	//	for i := 0; i < ref.numOfRows; i++ {
	//		for j := 0; j < ref.numOfColumns; j++ {
	//			v := ref.get(i, j)
	//			var groundType string
	//			var cost int
	//			if v.GroundData != nil {
	//				g := v.GroundData
	//				groundType = v.GroundData.ID
	//				cost = v.GroundData.Cost
	//			}
	//			fmt.Println(v.Position, cost, groundType)
	//		}
	//	}
	return ref.get(env.Start.Row, env.Start.Column), destinations, ref
}
