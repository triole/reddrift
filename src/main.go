package main

import (
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

	if CLI.Auto == true {
		loc := cap.GetLocation(CLI.Location)
		sp := suncalc.Init(loc.Coords.Lat, loc.Coords.Lon)
		tempInt = autoCalcTemp(sp.Altitude, CLI.Min, CLI.Max)
		lg.Log("Location is %+v\n", loc)
	}

	if CLI.Presets == true {
		listPresets()
		os.Exit(0)
	}

	if CLI.Repeat == true {
		for {
			loc := cap.GetLocation(CLI.Location)
			sp := suncalc.Init(loc.Coords.Lat, loc.Coords.Lon)
			tempInt = autoCalcTemp(sp.Altitude, CLI.Min, CLI.Max)
			setTemp(tempInt)
			time.Sleep(time.Duration(60) * time.Second)
		}
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
