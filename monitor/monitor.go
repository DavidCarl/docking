package monitor

import (
	"github.com/DavidCarl/docking/config"
	"github.com/DavidCarl/docking/logging"
)

// setupMonitors Needs to generate strings like this:
// xrandr --auto --output DVI-0 --mode 1440x900 --right-of DVI-1
func setupMonitors(modeInt int) {
	monitorName := make([]string, (len(screens) - 1))
	monitorResolution := make([]string, (len(screens) - 1))
	monitorActive := make([]bool, (len(screens) - 1))
	baseCommand := "xrandr -d :0.0 --auto --output "
	resolutionCommand := " --mode "

	for i := 0; i < (len(screens) - 1); i++ {
		monitorName[i], monitorResolution[i], monitorActive[i] = config.MonitorSetting(modeInt, screens[i])
		logging.Write(monitorName[i] + " " + monitorResolution[i])
	}

	for i := 0; i < (len(screens) - 1); i++ {
		if monitorActive[i] {
			xrandrCommand(baseCommand + monitorName[i] + resolutionCommand + monitorResolution[i])
		}
	}
}

func disableMonitors(modeInt int) {
	monitorName := make([]string, (len(screens) - 1))
	monitorResolution := make([]string, (len(screens) - 1))
	monitorActive := make([]bool, (len(screens) - 1))
	baseCommand := "xrandr -d :0.0 --output "
	disableCommand := " --off"
	disableMonitors := getMonitors(" disconnected")

	for i := 0; i < (len(screens) - 1); i++ {
		monitorName[i], monitorResolution[i], monitorActive[i] = config.MonitorSetting(modeInt, screens[i])
	}

	for i := 0; i < (len(disableMonitors) - 1); i++ {
		if disableMonitors[i] != "VIRTUAL1" {
			xrandrCommand(baseCommand + disableMonitors[i] + disableCommand)
		}
	}

	for i := 0; i < (len(screens) - 1); i++ {
		if !monitorActive[i] {
			xrandrCommand(baseCommand + monitorName[i] + disableCommand)
		}
	}
}
