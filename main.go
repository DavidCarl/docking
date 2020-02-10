package main

import (
	"fmt"
	"runtime"

	"github.com/DavidCarl/docking/logging"
	"github.com/DavidCarl/docking/monitor"
)

func main() {
	logging.Write("Starting docking application")
	if runtime.GOOS == "windows" {
		fmt.Println("Sorry, cant run this application on windows!")
	} else {
		monitor.Run()
	}
}
