package templeFights

import e "github.com/daniloanp/IA/environment"

type context struct {
	Saints  []e.Saint
	Temples []e.Temple
	MaxTime float64
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

//StringfyFighters ...
func StringfyFighters(f int, saints []e.Saint) []string {
	var str = make([]string, 0, len(saints))
	for inx, saint := range saints {
		if (1<<uint(inx))&f > 0 {
			str = append(str, saint.Name)
		}
	}

	return str
}

func decreaseLifes(fighters int, lifes []int) {
	for inx, _ := range lifes {
		if (1<<uint(inx))&fighters > 0 {
			lifes[inx]--
		}
	}
}

func encodeLiving(lifes []int) int {
	var living = int(0)

	for inx, lifes := range lifes {
		if lifes > 0 {
			living |= 1 << uint(inx)
		}
	}
	return living
}

func calcFightersPower(fighters int, saints []e.Saint) float64 {
	var sum = float64(0)

	for inx, saint := range saints {
		if (1<<uint(inx))&fighters > 0 {
			sum += saint.Power
		}
	}
	return sum
}

func calcCost(difficulty float64, fighters int, saints []e.Saint) float64 {
	var fightersPower = calcFightersPower(fighters, saints)
	return difficulty / fightersPower
}
