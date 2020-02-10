package monitor

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/DavidCarl/docking/config"
)

// Screens holds the screens in the dock
var screens []string

// Run This starts the chaos!
func Run() {
	screens = getConnectedMonitors()
	modeInt := config.FindMostLikelyConfig(screens)
	// fmt.Println(modeInt)
	setupMonitors(modeInt)
}

func getConnectedMonitors() []string {
	var output bytes.Buffer

	c1 := exec.Command("xrandr", "--query")
	c2 := exec.Command("grep", " connected")
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

// setupMonitors Needs to generate strings like this:
// xrandr --auto --output DVI-0 --mode 1440x900 --right-of DVI-1
func setupMonitors(modeInt int) {
	monitorName := make([]string, (len(screens) - 1))
	monitorResolution := make([]string, (len(screens) - 1))
	monitorActive := make([]bool, (len(screens) - 1))
	baseCommand := "xrandr --auto --output "
	resolutionCommand := " --mode "

	for i := 0; i < (len(screens) - 1); i++ {
		// fmt.Println(screens[i])
		monitorName[i], monitorResolution[i], monitorActive[i] = config.MonitorSetting(modeInt, screens[i])
	}

	for i := 0; i < (len(screens) - 1); i++ {
		if monitorActive[i] {
			fmt.Println(baseCommand + monitorName[i] + resolutionCommand + monitorResolution[i])
		}
	}
}
