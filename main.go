package main

import (
	"encoding/json" //importing(calling) library json
	"fmt"
	e "github.com/daniloanp/IA/environment"
	frontend "github.com/daniloanp/IA/frontend"
	walk "github.com/daniloanp/IA/pathThroughMap"
	fights "github.com/daniloanp/IA/templeFights"
	"io/ioutil"
	"time"
)

func main() {

	content, err := ioutil.ReadFile("default.map.json")

	if err != nil {
		fmt.Print("Error (reading file):", err)
		return
	}

	var conf e.Environment
	err = json.Unmarshal(content, &conf)

	if err != nil {
		fmt.Println("Error (parsing JSON)", err)
	}
	
	var initialTime = time.Now()
	origin, goals, _ := walk.BuildGraphFromEnvironment(&conf)
	var total int
	var paths = make([][]*walk.Square, 0, len(goals))
	for _, goal := range goals {
		res, duration := origin.AStar(goal)
		_, _ = res, duration
		total = duration + total

		fmt.Println("\n Inicio:")
		for i := len(res) - 1; i >= 0; i-- {
			square := res[i]
			fmt.Println("\t", square.Position)
		}
		fmt.Println("\nCusto para esse trajeto:", duration)
		paths = append(paths, res)
		origin = goal
	}

	//initAllegro(&conf)
	fmt.Println("\n\nTempo Total para andar no mapa:", total, "\n")
	fmt.Println("====================== x ======================")

	achou, res := fights.TemplesSolution(conf.Saints, conf.AvailableTime-float64(total), conf.Temples)

	var totalToFight = float64(0)
	if !achou {
		fmt.Println(":-(!")
		return
	}
	for i := len(res) - 1; i >= 0; i-- {
		var state = res[i]
		fmt.Println("\t Quem Lutou pra vir:", fights.StringfyFighters(state.Fighters, conf.Saints))
		fmt.Println("\t Vidas:", state.Lifes)
		fmt.Println("\t Tempo Gasto pra vir:", state.CostToMe(), "\n\n")
		totalToFight += state.CostToMe()
	}

	fmt.Println("Tempo Total para lutar:", totalToFight, "\n\n")
	fmt.Println("Tempo total pra salvar Saori:", totalToFight+float64(total))
	fmt.Println("Processing time: ", time.Now().Sub(initialTime))

	frontend.InitAllegro(&conf, paths, res)

}
