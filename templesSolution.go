package main

import (
	"fmt"
	"math"
)

type GameState struct {
	previousLifeOfSaints []int
	fighters             []int
	duration             int
	difficulty           int
	nextState            *GameState
}

var passoMaximo float32
var mapToPossibility MapToPossibility

func templesSolution(saints []Saint, avTime float32, temples []Temple) (bool, *GameState) {
	if mapToPossibility == nil {
		fmt.Println("\nGenerating...!\n")
		mapToPossibility.init(saints)
	}

	passoMaximo = avTime / float32(len(temples)+1)

	return backtrackedAStar(saints, avTime, temples)
}

func backtrackedAStar(saints []Saint, avTime float32, temples []Temple) (bool, *GameState) {
	var (
		lifeOfSaints = make([]int, 0, len(saints))
		listSaints   = make([]int, 0, len(saints))
		founded      = false
		current      *GameState
	)

	for inx, saint := range saints {
		lifeOfSaints = append(lifeOfSaints, saint.Lives)
		if saint.Lives > 0 {
			listSaints = append(listSaints, inx)
		}
	}

	//Ok!
	if len(listSaints) > 0 && len(temples) == 0 && avTime >= 0 {
		return true, nil
	}
	//Not OK!
	if len(listSaints) == 0 || len(temples) == 0 {
		fmt.Println("\n\ntodos mortos!\n\n")
		return false, nil
	}

	current = &GameState{
		previousLifeOfSaints: lifeOfSaints,
		difficulty:           temples[0].Difficulty,
	}

	var ref = mapToPossibility.get(listSaints)
	var res = make([][]int, len(ref))
	copy(res, ref)

	var wantedPower = float32(current.difficulty) / 32.2
	for !founded && len(res) > 0 {
		var minCost = float32(1 << 27)
		var minInx = -1
		var choosenFighters []int = nil
		// var minEstimatedCost = float32(1 << 27)
		var next *GameState
		// |res| , maybe a minHeap?

		var better = float32(1 << 27)

		for inx, fighters := range res {
			var cost = calcCost(current.difficulty, fighters, saints)
			var powerOfFighters = calcFightersPower(fighters, saints)
			if v := math.Abs(float64(wantedPower - powerOfFighters)); v < float64(better) {
				minInx = inx
				minCost = cost
				better = float32(v)
				choosenFighters = fighters
			}
		}

		if avTime-minCost < 0 {
			fmt.Println("\n\nsem tempo!\n\n")
			return false, nil
		}
		for _, inx := range choosenFighters {
			saints[inx].Lives--
		}

		founded, next = backtrackedAStar(saints, avTime-minCost, temples[1:])

		if !founded { // se não deu certo, restaura as vidas
			// remove already tried vertex from poss
			res = append(res[:minInx], res[1+minInx:]...)
			for _, inx := range choosenFighters {
				saints[inx].Lives++
			}
		} else { // se achou, gera vizinho
			current.duration = int(minCost)
			current.fighters = choosenFighters
			current.nextState = next
		}
		break
	}
	res = nil

	return founded, current
}

func heuristicForFighters(previousLifeOfSaints []int, fighters []int, saints []Saint, temples []Temple, cost float32) float32 {
	var numOfLiving = float64(len(previousLifeOfSaints))
	var numOfFighters = float64(len(fighters))
	var absOfLiving = math.Abs(numOfLiving/2-(numOfFighters)) * 10 //fmt.Println("Lutadores:", numOfFighters, "Gente viva:", numOfLiving, "Qualidade (menor é melhor):", absOfLiving)
	// numero de lutadores pequeno aumenta o custo, numero grande também
	var result = absOfLiving + math.Abs(float64(passoMaximo-cost))*2
	return float32(result)

}

func calcFightersPower(fighters []int, saints []Saint) float32 {
	var sum = float32(0)

	for _, fighter := range fighters {
		sum = float32(sum) + saints[fighter].Power
	}
	return sum
}

func calcCost(difficulty int, fighters []int, saints []Saint) float32 {
	var fightersPower = calcFightersPower(fighters, saints)
	return float32(difficulty) / fightersPower
}
