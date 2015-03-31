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
}
