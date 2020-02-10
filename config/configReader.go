package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
)

// Config This holds the base structure of our JSON object
type Config struct {
	Modes []Modes `json:"modes"`
}

// Modes This holds the different modes, docking, non-docking etc
type Modes struct {
	Monitors []Monitors `json:"monitors"`
	Name     string     `json:"name"`
}

// Monitors This holds information regarding the monitors in the modes
type Monitors struct {
	Name       string `json:"name"`
	Mode       bool   `json:"mode"`
	Resolution string `json:"resolution"`
	Position   string `json:"position"`
}

// loadJSONConfig This method is for loading and returning the JSON object
func loadJSONConfig() Config {
	usr, err := user.Current()
	file, err := os.Open(usr.HomeDir + "/.config/docking/config.json")
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println("Successfully Opened test.json")

	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var config Config

	json.Unmarshal(byteValue, &config)

	// for i := 0; i < len(config.Modes); i++ {
	// 	fmt.Println("Mode name:", config.Modes[i].Name)
	// 	for k := 0; k < len(config.Modes[i].Monitors); k++ {
	// 		fmt.Println("Monitor name:", config.Modes[i].Monitors[k].Name)
	// 		fmt.Println("Monitor mode:", config.Modes[i].Monitors[k].Mode)
	// 		fmt.Println("Monitor resolution:", config.Modes[i].Monitors[k].Resolution)
	// 	}
	// }

	return config
}

// FindMostLikelyConfig cyka
func FindMostLikelyConfig(sArray []string) int {
	var config Config
	var size int
	var mode int

	size = len(sArray) - 1

	config = loadJSONConfig()

	for i := 0; i < len(config.Modes); i++ {
		found := make([]bool, size)
		// fmt.Println("Mode name:", config.Modes[i].Name)
		for k := 0; k < len(config.Modes[i].Monitors); k++ {
			if len(config.Modes[i].Monitors) != size {
				break
			}
			mode = i
			for j := 0; j < size; j++ {
				// fmt.Println(sArray[j])
				if sArray[j] == config.Modes[i].Monitors[k].Name {
					found[j] = true
					// fmt.Println("Monitor name:", config.Modes[i].Monitors[k].Name)
					// fmt.Println("Monitor mode:", config.Modes[i].Monitors[k].Mode)
					// fmt.Println("Monitor resolution:", config.Modes[i].Monitors[k].Resolution)
				} else {
					found[j] = false
				}
			}
		}
		var foundAll bool = false
		for k := 0; k < size; k++ {
			if found[k] != true {
				break
			}
		}
		if foundAll {
			break
		}
	}

	return mode
}

// MonitorSetting this should setting values!
func MonitorSetting(mode int, screenName string) (string, string, bool) {
	var rtnName string
	var rtnResolution string
	var rtnMode bool
	config := loadJSONConfig()

	for range config.Modes {
		for index := range config.Modes[mode].Monitors {
			if config.Modes[mode].Monitors[index].Name == screenName {
				rtnName = config.Modes[mode].Monitors[index].Name
				rtnResolution = config.Modes[mode].Monitors[index].Resolution
				rtnMode = config.Modes[mode].Monitors[index].Mode
				break
			}
		}
	}

	// fmt.Println(rtnValue)
	return rtnName, rtnResolution, rtnMode
}
