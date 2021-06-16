package main

import (
	"os"
	"sct/capitals"
	log "sct/logging"
	"strconv"
	"time"
)

var (
	lg log.Logging
	ts TempSet
)

func main() {
	parseArgs()
	lg = log.Init(absPath(CLI.LogFile))

	ts.TempName = CLI.Temp
	ts.Temp, _ = strconv.Atoi(CLI.Temp)

	cap := capitals.Init()
	loc := cap.GetLocation(CLI.Location)

	ts.Capital = loc.Capital
	ts.Country = loc.Country
	ts.Lat = loc.Coords.Lat
	ts.Lon = loc.Coords.Lon

	if CLI.Repeat == true {
		c := time.Tick(time.Duration(CLI.TickInterval) * time.Second)
		for _ = range c {
			ts = autoAdjust(ts)
		}
		os.Exit(0)
	}

	if CLI.Auto == true {
		ts = autoAdjust(ts)
		os.Exit(0)
	}

	if CLI.Presets == true {
		listPresets()
		os.Exit(0)
	}

	// default action
	setTemp(ts)
}
