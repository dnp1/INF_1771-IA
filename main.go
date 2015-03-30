package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
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

func (m Map) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("     ")
	for i := int64(0); i < 2*int64(len(m.Base)); i = i + 1 {
		if (i % 20) == 0 {
			buffer.WriteString(strconv.FormatInt(i/2, 10))
		} else {
			buffer.WriteString(" ")
		}
	}
	buffer.WriteString("\n")

	for nl, l := range m.Base {
		aux := func(nl int) string {
			if nl >= 10 {
				return strconv.FormatInt(int64(nl), 10)
			}
			return " " + strconv.FormatInt(int64(nl), 10)

		}
		buffer.WriteString(aux(nl) + ": ")
		for _, c := range l {
			buffer.WriteString(" ")
			switch c {
			case "M":
				buffer.WriteString(" ")
			case "P":
				buffer.WriteString("~")
			case "R":
				buffer.WriteString("=")
			case "_":
				buffer.WriteString("T")
			case "S":
				buffer.WriteString("S")
			case "E":
				buffer.WriteString("E")
			default:
				log.Fatalln("Caracter Inválido: ", c)
			}

		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}

func main() {
	content, err := ioutil.ReadFile("default.map.json")

	if err != nil {
		fmt.Print("Error (reading file):", err)
		return
	}
	var conf Map
	err = json.Unmarshal(content, &conf)

	if err != nil {
		fmt.Println("Error (parsing JSON)", err)
	}

	fmt.Println(conf.String())
	// fmt.Println(conf.Base)
}
