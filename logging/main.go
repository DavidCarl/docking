package logging

import (
	"log"
	"log/syslog"
)

// Write call this function to write to the log!
func Write(message string) {
	logwriter, e := syslog.New(syslog.LOG_NOTICE, "docking")
	if e == nil {
		log.SetOutput(logwriter)
	}

	// Now from anywhere else in your program, you can use this:
	log.Print(message)
}
