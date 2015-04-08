package main

import "log"

//Matrix ..
type Matrix struct {
	data []*Square
	m    uint64
	n    uint64
}

//NewMatrix .
func NewMatrix(n, m uint64) *Matrix {
	return &Matrix{
		data: make([]*Square, n*m),
		m:    m,
		n:    n,
	}
}

const (
	TOP    = 0
	RIGHT  = 1
	BOTTOM = 2
	LEFT   = 3
)

func (mat *Matrix) get(i, j uint64) *Square {
	if i*mat.m+j < uint64(len(mat.data)) {
		return mat.data[i*mat.m+j]
	} else {
		log.Fatalln("Oops:", i*mat.m+j, i, j)
	}
	return nil
}

func (mat *Matrix) set(i uint64, j uint64, v *Square) {
	mat.data[i*mat.m+j] = v
}

func getCost(env *Environment, c string, x, y uint64) int64 {
	if c == "S" {
		return 0
	}
	for _, g := range env.Grounds {
		//if t.Position.X == uint64(x) && t.Position.Y == uint64(y) {
		if g.ID == c {
			return g.Cost
		}
	}
	// Temple?
	if c != "_" {
		log.Fatalln("oops: Invalid Square at", x, y, "!!!")
	}
	for _, t := range env.Temples {
		if t.Position.X == uint64(x) && t.Position.Y == uint64(y) {
			return t.Cost
		}
	}
	return 0
}

func getOrBuild(env *Environment, ref *Matrix, g string, x, y uint64) *Square {
	if v := ref.get(x, y); v != nil {
		return v
	}
	s := new(Square)
	s.Cost = getCost(env, g, x, y)
	s.p = Point{X: x, Y: y}
	s.neighbors = make([]*Square, 4)
	ref.set(x, y, s)
	return s
}

// buildGraphFromEnv build a Graph and return the initial and the final Square
func buildGraphFromEnv(env *Environment) (*Square, *Square) {
	ref := NewMatrix(42, 42)
	for x, l := range env.Map {
		for y, c := range l {
			var x, y = uint64(x), uint64(y)
			v := getOrBuild(env, ref, c, x, y)
			s := v
			if x > 0 { // has top neighbor
				s.neighbors[TOP] = getOrBuild(env, ref, c, x-1, y)
			}
			if y > 0 { // has  left neighbor
				s.neighbors[LEFT] = getOrBuild(env, ref, c, x, y-1)
			}
			if x < uint64(len(env.Map)-1) { // has bottom neighbor
				s.neighbors[BOTTOM] = getOrBuild(env, ref, c, x+1, y)
			}
			if y < uint64(len(l)-1) { // has right neighbor
				s.neighbors[RIGHT] = getOrBuild(env, ref, c, x, y+1)
			}
		}
	}
	return ref.get(env.Start.X, env.Start.Y), ref.get(env.End.X, env.End.Y)
}
