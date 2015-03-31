package main

import (
	"os"
	"os/exec"
	"runtime"
)

var clearWays map[string]func() //create a map for storing clear funcs

func init() {
	clearWays = make(map[string]func()) //Initialize it
	clearWays["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clearWays["windows"] = func() {
		cmd := exec.Command("cls") //Windows example it is untested, but I think its working
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// CallClear do something
func CallClear() {
	value, ok := clearWays[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                              //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

// func main() {
// 	fmt.Println("I will clean the screen in 2 seconds!")
// 	time.Sleep(2 * time.Second)
// 	CallClear()
// 	fmt.Println("I'm alone...")
// }
