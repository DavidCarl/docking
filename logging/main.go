package logging

import (
	"log"
	"os"
	"os/user"
)

// Write call this function to write to the log!
func Write(message string) {
	logState := false
	if logState {
		usr, _ := user.Current()
		file, err := os.OpenFile(usr.HomeDir+"/.config/docking/info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		log.SetOutput(file)
		log.Print(message)
	}
}
