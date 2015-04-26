package main

//MapToPossibility stores slices of possibilitites
type MapToPossibility map[int][][]int

func (m *MapToPossibility) get(key []int) [][]int {
	var inx = 0
	for _, v := range key {
		inx |= 1 << uint(v)
	}
	return (*m)[inx]
}

func (m *MapToPossibility) set(key []int, set [][]int) {
	var inx = 0
	for _, v := range key {
		inx |= 1 << uint(v)
	}
	(*m)[inx] = set
}

// generate all combinations and make cache of it
func (m *MapToPossibility) init(saints []Saint) {
	var listSaints = make([]int, len(saints))
	var S = make(map[int]bool)

	// Init Global state
	*m = MapToPossibility(make(map[int][][]int, intPow(2, len(saints))))

	// get an array of saints
	for inx, _ := range saints {
		listSaints[inx] = inx
	}
	// gera todas possiblidades de luta
	var res = make([][]int, 0, 31)
	for i := 1; i <= len(saints); i++ {
		combinations(&res, S, listSaints, i)
	}

	// for each subset of res, gen possible combination
	for _, living := range res {
		if len(living) == 1 {
			m.set(living, [][]int{living})
			continue
		} else if len(living) == len(saints) {
			m.set(living, res)
			continue
		}
		var possibilities = make([][]int, 0, intPow(2, len(living))-1)
		for i := 1; i <= len(living); i++ {
			combinations(&possibilities, S, living, i)
		}
		m.set(living, possibilities)
	}
}

func combinations(res *[][]int, S map[int]bool, l []int, k int) {
	if k == 0 { // Terminou
		if len(S) == 0 {
			return
		}
		var sl = make([]int, len(S)) // Criar um espaço pro resultado,
		var i = 0
		for v, _ := range S {
			sl[i] = v
			i++
		}
		*res = append(*res, sl)
		return
	}

	for j := 0; j < len(l)-(k-1); j++ {
		s := l[j]
		S[s] = true
		combinations(res, S, l[j+1:], k-1)
		delete(S, s)
	}
}

// auxiliary function to calc pow of integer (the math.Pow is for float)
func intPow(base, expoent int) int {
	var accum = 1
	for ; expoent > 0; expoent-- {
		accum *= base
	}
	return accum
}
