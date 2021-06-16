package main

import (
	"fmt"
	"strconv"
	"time"

	"reddrift/suncalc"
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

func updateValues(ts tempSet, t time.Time) (r tempSet) {
	r = ts
	sc := suncalc.Init(ts.Lat, ts.Lon, t)
	r.SunAltitude = sc.Altitude
	r.SunAzimuth = sc.Azimuth
	r.LastTemp = readStatusFile()
	r.Temp = autoCalcTemp(r, CLI.TempMin, CLI.TempMax)
	return
}

func autoCalcTemp(ts tempSet, min, max int) (temp int) {
	temp = min
	diff := float64(max-min) * (ts.SunAltitude + CLI.Postpone)
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
		if ts.Temp > CLI.TempMax {
			ts.Temp = CLI.TempMax
		}
		if ts.Temp < CLI.TempMin {
			ts.Temp = CLI.TempMin
		}
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

func showDayCycle(ts tempSet) {
	fmt.Printf("\n%s %f %f\n\n", ts.Capital, ts.Lat, ts.Lon)
	t := time.Now()
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, t.Nanosecond(), t.Location())
	for i := 0; i <= 47; i++ {
		ts = updateValues(ts, t)
		fmt.Printf("%s\t%+f\t%d\n", t.Format("2006-01-02 15:04"), ts.SunAltitude, ts.Temp)
		t = t.Add(time.Duration(30) * time.Minute)
	}
	fmt.Println("")
}
