package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("default.map.json")

	if err != nil {
		fmt.Print("Error (reading file):", err)
		return
	}

	var conf Environment
	err = json.Unmarshal(content, &conf)

	if err != nil {
		fmt.Println("Error (parsing JSON)", err)
	}

	v1, v2 := buildGraphFromEnv(&conf)
	fmt.Println(v1.BFS(v2))
}
