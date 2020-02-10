package main

import (
	"fmt"
	"github.com/DavidCarl/docking/monitor"
	"runtime"
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Sorry, cant run this application on windows!")
	} else {
		monitor.Run()
	}
}
