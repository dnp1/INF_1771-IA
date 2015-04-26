package main

import (
	"fmt"
	"math"
)

var passoMaximo float64
var mapToPossibility MapToPossibility
var Temples []Temple

var Lives = make(map[GameState][]int)

func templesSolution(saints []Saint, avTime float64, temples []Temple) (bool, []*GameState) {
	if mapToPossibility == nil {
		fmt.Println("\nGenerating...!\n")
		mapToPossibility.init(saints)
	}

	passoMaximo = avTime / float64(len(temples)+1)
	Temples = temples

	return backtrackedAStar(saints, avTime, 0)
}

func backtrackedAStar(saints []Saint, avTime float64, inxTemple int) (bool, []*GameState) {
	var (
		livingSaints = int(0)
		listSaints   = make([]int, 0, len(saints))
		founded      = false
		current      *GameState
		result       []*GameState
	)

	for inx, saint := range saints {
		if saint.Lives > 0 {
			livingSaints |= 1 << uint(inx)
			listSaints = append(listSaints, inx)
		}
	}

	//Ok!
	if livingSaints > 0 && inxTemple == len(Temples) && avTime >= 0 {
		lives := make([]int, len(saints))
		for inx, saint := range saints {
			lives[inx] = saint.Lives
		}
		current = &GameState{
			living:    livingSaints,
			inxTemple: inxTemple,
			timeLeft:  avTime,
		}
		Lives[*current] = lives
		return true, []*GameState{current}
	}
	//Not OK!
	if livingSaints == 0 || inxTemple == len(Temples) {
		fmt.Println("\n\ntodos mortos!\n\n")
		return false, nil
	}

	current = &GameState{
		living:    livingSaints,
		inxTemple: inxTemple,
		timeLeft:  avTime,
	}

	var res = mapToPossibility.get(listSaints)

	var difficulty = Temples[inxTemple].Difficulty
	var wantedPower = difficulty / 32.2
	for !founded && len(res) > 0 {
		var minCost = float64(1 << 27)
		var better = float64(1 << 27)
		var minInx = -1
		var choosenFighters []int = nil

		for inx, fighters := range res {
			var cost = calcCost(difficulty, fighters, saints)
			var powerOfFighters = calcFightersPower(fighters, saints)
			if v := math.Abs(wantedPower - powerOfFighters); v < better {
				minInx = inx
				minCost = cost
				better = v
				choosenFighters = fighters
			}
		}
		fmt.Println("chosenFighters:", choosenFighters)
		if avTime-minCost < 0 {
			fmt.Println("\n\nsem tempo!\n\n")
			return false, nil
		}
		for _, inx := range choosenFighters {
			saints[inx].Lives--
		}

		avTime = math.Trunc((avTime-minCost)*1000) / 1000
		founded, result = backtrackedAStar(saints, avTime, inxTemple+1)

		for _, inx := range choosenFighters {
			saints[inx].Lives++
		}

		if !founded { // se não deu certo, restaura as vidas
			// remove already tried vertex from poss
			res = append(res[:minInx], res[1+minInx:]...)

		} else {
			lives := make([]int, len(saints))
			for inx, saint := range saints {
				lives[inx] = saint.Lives
			}

			Lives[*current] = lives
			return true, append([]*GameState{current}, result...)
		}
		break
	}
	res = nil

	return false, nil
}

//func heuristicForFighters(previousLifeOfSaints []int, fighters []int, saints []Saint, temples []Temple, cost float32) float32 {
//	var numOfLiving = float64(len(previousLifeOfSaints))
//	var numOfFighters = float64(len(fighters))
//	var absOfLiving = math.Abs(numOfLiving/2-(numOfFighters)) * 10 //fmt.Println("Lutadores:", numOfFighters, "Gente viva:", numOfLiving, "Qualidade
//	var result = absOfLiving + math.Abs(float64(passoMaximo-cost))*2
//	return float32(result)

//}

func calcFightersPower(fighters []int, saints []Saint) float64 {
	var sum = float64(0)

	for _, fighter := range fighters {
		sum = sum + saints[fighter].Power
	}
	return sum
}

func calcCost(difficulty float64, fighters []int, saints []Saint) float64 {
	var fightersPower = calcFightersPower(fighters, saints)
	return difficulty / fightersPower
}
