package main

import (
	"encoding/json" //importing(calling) library json
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("default.map.json")
	//content-> receive the content from default.map.json
	//err -> receive condRet
	if err != nil {
		fmt.Print("Error (reading file):", err)
		return
	}

	var conf Environment
	err = json.Unmarshal(content, &conf)

	if err != nil {
		fmt.Println("Error (parsing JSON)", err)
	}

	origin, goals, _ := buildGraphFromEnvironment(&conf)
	var total int
	for _, goal := range goals {
		res, duration := origin.AStar(goal)
		_, _ = res, duration
		total = duration + total
		origin = goal
	}

	//initAllegro(&conf)
	//	fmt.Println("\n\nTotal Duration:", total, "\n")

	achou, resultado := templesSolution(conf.Saints, float32(conf.AvailableTime-total), conf.Temples)

	if !achou {
		fmt.Println(":-(!")
		return
	}
	for ; resultado != nil; resultado = resultado.nextState {
		fmt.Println(resultado)
	}

}
