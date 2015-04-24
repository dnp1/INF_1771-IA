package main

import (
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

//note: the sintax `json:"id"` bind the identifier in json to the variable in this code
//example: X uint64 `json:"x"` bind the name the x name in json to X point struct

// Point define a 2-axis coordinate type
type Point struct {
	Row    uint64 `json:"row"`
	Column uint64 `json:"column"`
}

// Ground is where the characters go by in the map. It has a fixed cost to pass through it
type Ground struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Cost int64  `json:"cost"`
}

// Temple is a special place in the map with fixed cost to pass through it,
// Code irrelevante -> the cost are little bite higher than the ground because the temple keep the Gold Knights.
type Temple struct {
	Name      string `json:"name"`
	Dificulty int64  `json:"cost"`
	Position  Point  `json:"position"`
}

// Environment of the game
type Environment struct {
	Start   Point      `json:"start"`
	End     Point      `json:"end"`
	Grounds []Ground   `json:"grounds"`
	Temples []Temple   `json:"temples"`
	Map     [][]string `json:"map"`
	printed bool
}

func (m Environment) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("     ")
	for i := int64(0); i < 2*int64(len(m.Map)); i = i + 1 {
		if (i % 20) == 0 {
			buffer.WriteString(strconv.FormatInt(i/2, 10))
		} else {
			buffer.WriteString(" ")
		}
	}
	buffer.WriteString("\r\n")

	aux := func(numlines int) string { //the objective here is print the row numbers aligned
		if numlines >= 10 { //if true, will print the number higher than 10
			return strconv.FormatInt(int64(numlines), 10)

		} //else, will print a space plus a number lower than 10
		return " " + strconv.FormatInt(int64(numlines), 10)
	}

	for numlines, lines := range m.Map {

		buffer.WriteString(aux(numlines) + ": ")
		for _, c := range lines { // print each simbol in map(m.Base)
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
			case "S": //S is not defined
				buffer.WriteString("S")
			case "E": // E is not defined
				buffer.WriteString("E")
			default:
				log.Fatalln("Caracter Inválido: ", c)
			}
		}
		buffer.WriteString("\r\n")
	}
	return buffer.String()
}

func clear(amount uint64) {
	for i := uint64(0); i < amount; i++ {
		os.Stdout.WriteString("\033[A\033[2K")
	}
	// what those below do?
	os.Stdout.Seek(0, 0)
	os.Stdout.Truncate(0) /* you probably want this as well */
	os.Stdout.Sync()

}

//
// Print the map
func (m *Environment) Print() {
	dat := []byte(m.String())

	i := uint64(0)

	// work around to get line number
	strings.Map(func(r rune) rune {
		if r == '\n' {
			i = i + 1
		}
		return r
	}, m.String())

	if m.printed {
		clear(uint64(i))
	} else {
		m.printed = true
	}
	os.Stdout.Write(dat)

	os.Stdout.Sync()
	time.Sleep(2 << 31)
}
