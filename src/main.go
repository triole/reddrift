package main

import (
	"math"
	"os"
	"reddrift/capitals"
	log "reddrift/logging"
	"strconv"
	"time"
)

var (
	lg log.Logging
	ts tempSet
)

func main() {
	var err error
	parseArgs()
	lg = log.Init(absPath(CLI.LogFile))

	ts.TempName = CLI.Temp
	if ts.TempName == "6500" {
		ts.TempName = "default"
	}
	ts.Temp, _ = strconv.Atoi(CLI.Temp)

	cap := capitals.Init()
	loc := cap.GetLocation(CLI.Location)

	ts.Capital = loc.Capital
	ts.Country = loc.Country
	ts.Lat = loc.Coords.Lat
	ts.Lon = loc.Coords.Lon

	ts.CurveModifier = CLI.CurveModifier
	ts.CurveModList = makeCurveModList()

	ts.ValidPreset = false
	ts.ValidTempInt = false

	if CLI.ListPresets == true {
		listPresets()
		os.Exit(0)
	}

	if CLI.ListLocations == true {
		cap.ListLocations()
		os.Exit(0)
	}

	if CLI.DayCycle == true {
		showDayCycle(ts)
		os.Exit(0)
	}

	// default action
	ts = updateValues(ts, time.Now())
	if ts.TempName != "default" {
		temp := Temp{}
		temp, err = temp.set(ts.TempName)
		if err == nil {
			ts.Temp = temp.Value
		}
	}
	setTemp(ts)

	if CLI.Repeat == true {
		c := time.Tick(time.Duration(CLI.TickInterval) * time.Second)
		for _ = range c {
			ts = updateValues(ts, time.Now())
			setTemp(ts)
		}
	}
}

func makeCurveModList() (fl []float64) {
	frag := ts.CurveModifier / 4
	fl = append(fl, round(ts.CurveModifier-(2*frag)))
	fl = append(fl, round(ts.CurveModifier-frag))
	fl = append(fl, round(ts.CurveModifier))
	fl = append(fl, round(ts.CurveModifier+frag))
	fl = append(fl, round(ts.CurveModifier+(2*frag)))
	return
}

func round(x float64) float64 {
	return math.Round(x*100) / 100
}
