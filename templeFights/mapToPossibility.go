package templeFights

import e "github.com/daniloanp/IA/environment"

//MapToPossibility stores slices of possibilitites (encoded as int)
type MapToPossibility []([]int)

func (m *MapToPossibility) get(key int) []int {
	return (*m)[key]
}

func (m *MapToPossibility) set(key int, set []int) {
	(*m)[key] = set
}

// generate all combinations and make cache of it
func (m *MapToPossibility) init(saints []e.Saint) {
	var intSaints int

	// Init Global state
	*m = MapToPossibility(make([][]int, intPow(2, len(saints))))
	// get an array of saints
	for inx, _ := range saints {
		intSaints |= (1 << uint(inx))
	}
	// gera todas possiblidades de luta
	var res = make([]int, 0, 31)
	for i := 1; i <= len(saints); i++ {
		combinations(&res, 0, intSaints, i, 0)
	}

	// for each subset of res, gen possible combination
	for _, living := range res {

		var n = int(0)
		for i := uint(0); (1 << i) < living; i++ {
			if (1<<i)&living > 0 {
				n++
			}
		}

		if n == 1 {
			m.set(living, []int{living})
			continue
		} else if n+1 == intPow(2, len(saints)) {
			m.set(living, res)
			continue
		}
		var possibilities = make([]int, 0, n)
		for i := 1; i <= n; i++ {
			combinations(&possibilities, 0, living, i, 0)
		}
		m.set(living, possibilities)
	}
}
