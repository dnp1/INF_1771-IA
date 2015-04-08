package main

import (
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
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

// Line of a map is just an slice of Cells
type Line []string

// Environment of the game
type Environment struct {
	Start   Point    `json:"start"`
	End     Point    `json:"end"`
	Grounds []Ground `json:"grounds"`
	Temples []Temple `json:"temples"`
	Map     []Line   `json:"map"`
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

	aux := func(nl int) string {
		if nl >= 10 {
			return strconv.FormatInt(int64(nl), 10)
		}
		return " " + strconv.FormatInt(int64(nl), 10)
	}

	for nl, l := range m.Map {
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
		buffer.WriteString("\r\n")
	}
	return buffer.String()
}

func clear(amount uint64) {
	for i := uint64(0); i < amount; i++ {
		os.Stdout.WriteString("\033[A\033[2K")
	}
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
