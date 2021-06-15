package main

import (
	"fmt"
	"os"
	"sct/capitals"
	log "sct/logging"
	"sct/suncalc"
	"strconv"
	"time"
)

var (
	lg log.Logging
)

func main() {
	parseArgs()
	lg = log.Init(absPath(CLI.LogFile))

	tempInt := CLI.Temp
	cap := capitals.Init()
	loc := cap.GetLocation(CLI.Location)

	if CLI.Repeat == true {
		c := time.Tick(time.Duration(CLI.TickInterval) * time.Second)
		for _ = range c {
			autoAdjust(loc.Coords.Lat, loc.Coords.Lon, fmt.Sprintf("%+v", loc))
		}
		os.Exit(0)
	}

	if CLI.Auto == true {
		autoAdjust(loc.Coords.Lat, loc.Coords.Lon, fmt.Sprintf("%+v", loc))
		os.Exit(0)
	}

	if CLI.Presets == true {
		listPresets()
		os.Exit(0)
	}

	// default action
	setTemp(tempInt)
}

func autoCalcTemp(altitude float64, min, max int) (s string) {
	temp := min
	diff := float64(max-min) * (altitude)
	if diff >= 0 && diff < 1 {
		temp = min + int(diff)
	}
	if diff >= 1 {
		temp = max
	}
	s = strconv.Itoa(temp)
	return
}

func autoAdjust(lat, lon float64, loc string) (tempInt string, didChange bool) {
	sd := suncalc.Init(lat, lon)
	tempInt = autoCalcTemp(sd.Altitude, CLI.Min, CLI.Max)
	didChange = readStatusFile() != tempInt
	if didChange == true {
		lg.Log("Location is %q\n", loc)
		setTemp(tempInt)
	}
	return
}
