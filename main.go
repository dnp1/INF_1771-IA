package main

import (
	"encoding/json" //importing(calling) library json
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("default.map.json")
	//content-> receive the content from default.map.json
	//err -> receive assert
	if err != nil {
		fmt.Print("Error (reading file):", err)
		return
	}

	var conf Environment
	err = json.Unmarshal(content, &conf)

	if err != nil {
		fmt.Println("Error (parsing JSON)", err)
	}

	origin, goals := buildGraphFromEnv(&conf)

	for _, goal := range goals {
		res := origin.AStar(goal)
		for _, v := range res {
			fmt.Print(v.Position, " ,")
		}
		fmt.Print("\n")
		origin = goal
	}
}
