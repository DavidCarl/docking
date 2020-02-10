package monitor

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/DavidCarl/docking/config"
	"github.com/DavidCarl/docking/logging"
)

// Screens holds the screens in the dock
var screens []string

// Run This starts the chaos!
func Run() {
	screens = getMonitors(" connected")
	modeInt := config.FindMostLikelyConfig(screens)
	setupMonitors(modeInt)
	disableMonitors(modeInt)
}

func getMonitors(state string) []string {
	logging.Write("getMonitors " + state)
	var output bytes.Buffer

	c1 := exec.Command("xrandr", "-d", ":0.0", "--query")
	c2 := exec.Command("grep", state)
	c3 := exec.Command("cut", "-d ", "-f1")
	c2.Stdin, _ = c1.StdoutPipe()
	c3.Stdin, _ = c2.StdoutPipe()
	c3.Stdout = &output // Put Stdout from c3 into the bytes.Buffer (output)
	_ = c3.Start()
	_ = c2.Start()
	_ = c1.Run()
	_ = c2.Wait()
	_ = c3.Wait()

	return strings.Split(output.String(), "\n")
}

func xrandrCommand(command string) bool {
	logging.Write("xrandrCommand " + command)
	// fmt.Println(command)
	parts := strings.Split(command, " ")
	head := parts[0]
	args := parts[1:len(parts)]
	cmd := exec.Command(head, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return false
	}
	return true
}
