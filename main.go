package main

import (
	"fmt"
	"runtime"

	"github.com/DavidCarl/docking/logging"
	"github.com/DavidCarl/docking/monitor"
)

func main() {
	// command := os.Args[1]
	// fmt.Println(command)
	// if command == "setup" {
	// 	path := "/var/log/docking"
	// 	if _, err := os.Stat(path); os.IsNotExist(err) {

	// 		os.Mkdir(path, 0700)
	// 	}
	// } else {
	logging.Write("Starting docking application")
	if runtime.GOOS == "windows" {
		fmt.Println("Sorry, cant run this application on windows!")
	} else {
		monitor.Run()
	}
	// }
}
