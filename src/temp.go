package main

import (
	"fmt"
	"strconv"

	"github.com/dim13/sct/src/suncalc"
)

type Temp struct {
	Value int
	Name  string
}

func (m *Temp) Set(s string) (t Temp, err error) {
	var i int
	ts.ValidPreset = false
	ts.ValidTempInt = false
	if val, ok := presets[s]; ok {
		t = Temp{
			Name:  s,
			Value: val,
		}
		ts.ValidPreset = true
	} else {
		i, _ = strconv.Atoi(s)
		if i < presets["minimal"] || i > presets["maximal"] {
			return t, fmt.Errorf(
				"Temp out of range, value needs to be between %v and %v",
				presets["minimal"], presets["maximal"],
			)
		}
		t.Name = "manual"
		t.Value = i
		ts.ValidTempInt = true
	}
	return
}

func updateValues(ts tempSet) (r tempSet) {
	r = ts
	sc := suncalc.Init(ts.Lat, ts.Lon)
	r.SunAltitude = sc.Altitude
	r.SunAzimuth = sc.Azimuth
	r.LastTemp = readStatusFile()
	r.Temp = autoCalcTemp(r, CLI.Min, CLI.Max)
	return
}

func autoCalcTemp(ts tempSet, min, max int) (temp int) {
	temp = min
	al := ts.SunAltitude + 0.3
	diff := float64(max-min) * (al)
	temp = min + int(diff)
	if temp < min {
		temp = min
	}
	if temp > max {
		temp = max
	}
	ts.Temp = temp
	return
}

func setTemp(ts tempSet) {
	if ts.Temp != ts.LastTemp || CLI.Force == true {
		if ts.ValidPreset == true {
			lg.Log("Set to temp to \""+ts.TempName+"\" %d\n", ts.Temp)
		}
		if ts.ValidTempInt == true {
			lg.Log("Set to temp %d\n", ts.Temp)
		}
		if ts.ValidPreset == false && ts.ValidTempInt == false {
			lg.Log("Set to temp %+v\n", ts)
		}
		Set(ts.Temp)
		saveStatusFile(strconv.Itoa(ts.Temp))
	}
}
