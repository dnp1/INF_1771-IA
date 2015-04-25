package main

import "fmt"

import "log"
import "math"

type GameState struct {
	previousLifeOfSaints []int
	fighters             []int
	duration             int
	difficulty           int
	nextState            *GameState
}

const DEBUG = true

func myPrintln(l ...interface{}) {
	if DEBUG {
		fmt.Println(l...)
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

//QUalidade da escolha
func heuristicForFighters(previousLifeOfSaints []int, fighters []int, saints []Saint, temples []Temple) float32 {
	var numOfLiving = float64(len(previousLifeOfSaints))
	var numOfFighters = float64(len(fighters))
	var absOfLiving = math.Abs(numOfLiving/2-(numOfFighters-1)) * 5
	//fmt.Println("Lutadores:", numOfFighters, "Gente viva:", numOfLiving, "Qualidade (menor é melhor):", absOfLiving)
	// numero de lutadores pequeno aumenta o custo, numero grande também
	return float32(absOfLiving)

}

func calcCost(difficulty int, fighters []int, saints []Saint) float32 {
	var sum = float32(0)

	for _, fighter := range fighters {
		sum = float32(sum) + saints[fighter].Power
	}

	return float32(difficulty) / sum
}

func buildGraph(saints []Saint, avTime int, temples []Temple, depth int) (bool, *GameState) {
	var (
		lifeOfSaints = make([]int, 0, len(saints))
		listSaints   = make([]int, 0, len(saints))
		res          = make([][]int, 0, 31)
		S            = make(map[int]bool)
		founded      = false
	)

	var l log.Logger
	_ = l

	for inx, saint := range saints {
		lifeOfSaints = append(lifeOfSaints, saint.Lives)
		if saint.Lives > 0 {
			listSaints = append(listSaints, inx)
		}
	}
	// Ok!
	if len(listSaints) > 0 && len(temples) == 0 && avTime >= 0 {
		fmt.Println("ACHOU!: ", avTime)
		return true, nil // founded
	}

	// Not Ok!
	if len(listSaints) == 0 || avTime < 0 {

		return false, nil
	}
	var initial = &GameState{
		previousLifeOfSaints: lifeOfSaints,
		difficulty:           temples[0].Difficulty,
	}

	// gera todas possiblidades de luta
	for i := 1; i <= len(saints); i++ {
		combinations(&res, S, listSaints, i)
	}

	// decide por onde seguir

	var initialLen = len(res)
	for !founded && len(res) > 0 {
		var minCost = 1 << 27
		var minInx = -1
		var choosenFighters []int = nil
		var minEstimatedCost = 1 << 27
		var next *GameState
		// |res| , maybe a minHeap?
		for inx, fighters := range res {

			var cost = calcCost(initial.difficulty, fighters, saints)
			var estimatedCost = cost + heuristicForFighters(listSaints, fighters, saints, temples)

			// decreasing lives
			if int(estimatedCost) < minEstimatedCost {
				minCost = int(cost)
				minInx = inx // Quem vai lutar
				minEstimatedCost = int(estimatedCost)
				choosenFighters = fighters
			}
		}

		if avTime-int(minCost) < 0 {
			return false, nil
		}

		fmt.Println("vidas:", lifeOfSaints)
		fmt.Println("Prox. Lutadores:", choosenFighters)
		fmt.Println("Tempo Disponível:", avTime)
		fmt.Println("Custo Prox. Mov:", minCost)
		fmt.Println("Profundidade:", depth)
		fmt.Println("\n")

		for _, inx := range choosenFighters {
			saints[inx].Lives--
		}

		// remove already tried vertex from poss
		res = append(res[:minInx], res[1+minInx:]...)
		if len(res) >= initialLen {
			log.Fatalln()
		}

		// Tenta seguir, se não for possível, volta

		founded, next = buildGraph(saints, avTime-int(minCost), temples[1:], depth+1)
		if !founded { // não deu certo, restaura as vidas
			for _, inx := range choosenFighters {
				saints[inx].Lives++
			}
		} else {
			initial.duration = int(minCost)
			initial.fighters = choosenFighters
			initial.nextState = next
		}
	}
	res = nil
	S = nil

	return founded, initial
	//	myPrintln(len(res), "\n", initial)
}
