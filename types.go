package main

import (
	"bytes"
	"log"
	"os"
	"strconv"
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

func clear(amount uint64) {
	for i := uint64(0); i < amount; i++ {
		os.Stdout.WriteString("\033[H\033[2J")
	}
	os.Stdout.Sync()
}

// Print the map
func (m Map) Print() {

	// i, err :=

	dat := []byte(m.String())
	if m.printed {
		clear(uint64(len(dat)))
	}
	os.Stdout.Write(dat)
	// // os.Stdout.Seek(0, 0)
	// if i, err := os.Stdout.WriteAt(dat, int64(len(dat))); err != nil {
	// 	log.Println(i)
	// 	log.Fatalln(err)
	// } else {
	// 	fmt.Println(i)
	// }
	os.Stdout.Sync()
	time.Sleep(2 << 32)

}
