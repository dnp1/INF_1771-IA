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
	X uint64 `json:"x"` 
	Y uint64 `json:"y"`
}

// Ground is where the characters go by in the map. it has a fixed cost to pass through it
type Ground struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Cost int64  `json:"cost"`
}

// Temple is a special place in the map with fixed cost to pass through it, the cost are little bite higher than the ground because the temple keep the Gold Knights. 
type Temple struct {
	Name     string `json:"name"`
	Cost     int64  `json:"cost"`
	Position Point  `json:"position"`
}

// Line of a map is just an slice of Cells
type Line []string

// Map of the game
type Map struct {
	Start   Point    `json:"start"`
	End     Point    `json:"end"`
	Grounds []Ground `json:"grounds"`
	Temples []Temple `json:"temples"`
	Base    []Line   `json:"base"`
	printed bool
}

func (m Map) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("     ")
	for i := int64(0); i < 2*int64(len(m.Base)); i = i + 1 {// here will print the number of the coluns
		if (i % 20) == 0 {
			buffer.WriteString(strconv.FormatInt(i/2, 10))
		} else {
			buffer.WriteString(" ")
		}
	}
	buffer.WriteString("\r\n")

	for numlines, lines := range m.Base {
		aux := func(numlines int) string {//the objective here is print the row numbers aligned 
			if numlines >= 10 {//if true, will print the number higher than 10 
				return strconv.FormatInt(int64(numlines), 10)
			}//else, will print a space plus a number lower than 10
			return " " + strconv.FormatInt(int64(numlines), 10)
		
		}
		buffer.WriteString(aux(numlines) + ": ")
		for _, c := range lines {// print each simbol in map(m.Base) 
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

func clear(amount uint64) {
	for i := uint64(0); i < amount; i++ {
		os.Stdout.WriteString("\033[A\033[2K")
	}
	// what those below do?
	os.Stdout.Seek(0, 0)
	os.Stdout.Truncate(0) /* you probably want this as well */
	os.Stdout.Sync()

}

// Print the map
func (m *Map) Print() {
	dat := []byte(m.String())

	i := uint64(0)

	// work around to get all
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
