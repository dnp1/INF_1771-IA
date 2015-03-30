package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Point define a 2-axis coordinate type
type Point struct {
	X uint64 `json:"x"`
	Y uint64 `json:"y"`
}

// Ground is a kind of space in the map with fixed cost
type Ground struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Cost int64  `json:"cost"`
}

// Temple is a special space in the map with fixed cost
type Temple struct {
	Name     string `json:"name"`
	Cost     int64  `json:"cost"`
	Position Point  `json:"position"`
}

// Line of a map
type Line []string

// Map of the game
type Map struct {
	Start   Point    `json:"start"`
	End     Point    `json:"end"`
	Grounds []Ground `json:"grounds"`
	Temples []Temple `json:"temples"`
	Base    []Line   `json:"base"`
}

func main() {
	content, err := ioutil.ReadFile("default.map.json")
	//fmt.Println(content)
	// importMap()
	if err != nil {
		fmt.Print("Error:", err)
	}
	var conf Map
	err = json.Unmarshal(content, &conf)
	if err != nil {
		fmt.Println("PUTS")
		fmt.Println(err)
	}

	fmt.Println(conf)
}
