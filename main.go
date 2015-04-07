package main

import (
	"encoding/json"//importing(calling) library json
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
	var conf Map
	err = json.Unmarshal(content, &conf)

	if err != nil {
		fmt.Println("Error (parsing JSON)", err)
	}

	conf.Print()
	//
	conf.Base[1][12] = "_"

	conf.Print()

	// fmt.Println(conf.Base)

	// os.Stdout.WriteString("output1\n")
	// os.Stdout.WriteString("output2\n")
	// time.Sleep(2 << 28)
	// os.Stdout.Sync()
	// os.Stdout.WriteString("\033[A\033[2K\033[A\033[2K")
	//
	// os.Stdout.Seek(0, 0)
	// os.Stdout.Truncate(1) /* you probably want this as well */
	// os.Stdout.Sync()
	//
	// os.Stdout.WriteString("output3\n")
	// os.Stdout.WriteString("output4\n")
	// os.Stdout.Sync()
}
