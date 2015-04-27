package templeFights

import e "github.com/daniloanp/IA/environment"

//MapToPossibility stores slices of possibilitites
type MapToPossibility []([]int)

func (m *MapToPossibility) get(key int) []int {
	return (*m)[key]
}

func (m *MapToPossibility) set(key int, set []int) {
	(*m)[key] = set
}

// generate all combinations and make cache of it
func (m *MapToPossibility) init(saints []e.Saint) {
	//var listSaints = make([]int, len(saints))
	var intSaints int
	//var S int

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

func combinations(res *[]int, S int, l int, k int, base int) {
	if k == 0 { // Terminou
		if S == 0 {
			return
		}
		*res = append(*res, S)
		return
	}
	var n = int(0)

	for i := uint(0); (1 << i) < l+1; i++ {
		if (1<<i)&l > 0 {
			n++
		}
	}

	//fmt.Println("LEN:", n)

	for j, visited := 0, 0; visited < n-(k-1); j++ {
		if s := (1 << uint(j)) & l; s > 0 {
			visited++
			S |= (1 << uint(j+base))
			//S[s] = true
			combinations(res, S, (l >> uint(j+1)), k-1, base+j+1)

			S ^= (1 << uint(j+base))
		}
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

func StringfyFighters(f int, saints []e.Saint) string {
	var str = "\n\t\t{\n"
	for inx, saint := range saints {

		if (1<<uint(inx))&f > 0 {
			str += "\t\t" + saint.Name + "\n"
		}

	}
	str += "\t\t}"
	return str
}
