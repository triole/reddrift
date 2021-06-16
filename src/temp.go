package main

import (
	"fmt"
	"strconv"

	"github.com/dim13/sct/src/suncalc"
)

var (
	presets = map[string]int{
		"minimal":     1000,
		"candle":      2300,
		"tungsten":    2700,
		"halogen":     3400,
		"fluorescent": 4200,
		"daylight":    5000,
		"default":     6500,
		"upallnight":  8000,
		"maximal":     10000,
	}
)

type Temp struct {
	Value int
	Name  string
}

func (m *Temp) Set(s string) (t Temp, err error) {
	if val, ok := presets[s]; ok {
		t = Temp{
			Name:  s,
			Value: val,
		}
	} else {
		i, err := strconv.Atoi(s)
		lg.LogIfErr(err, "Can not parse string to int %d", s)
		if i < presets["minimal"] || i > presets["maximal"] {
			return t, fmt.Errorf(
				"Temp out of range, value needs to be between %v and %v",
				presets["minimal"], presets["maximal"],
			)
		}
		t.Name = "manual"
		t.Value = i
	}
	return
}

func listPresets() {
	if CLI.Presets == true {
		fmt.Println("\nAvailable presets")
		hackedMap, hackedKeys := sortMapHack(presets)
		for _, k := range hackedKeys {
			fmt.Printf("%v\t%s\n", k, hackedMap[k])
		}
		fmt.Printf("\n")
	}

}

func autoCalcTemp(ts TempSet, min, max int) (temp int) {
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
	return
}

func autoAdjust(ts TempSet) TempSet {
	sc := suncalc.Init(ts.Lat, ts.Lon)

	ts.SunAltitude = sc.Altitude
	ts.SunAzimuth = sc.Azimuth
	ts.LastTemp = readStatusFile()
	ts.Temp = autoCalcTemp(ts, CLI.Min, CLI.Max)

	if ts.Temp != ts.LastTemp || CLI.Force == true {
		setTemp(ts)
	}
	return ts
}

func setTemp(ts TempSet) {
	temp := Temp{}
	temp, err := temp.Set(ts.TempName)
	ts.Temp = temp.Value
	if err != nil {
		fmt.Printf("%q\n", err)
	} else {
		lg.Log("Set to %+v\n", ts)
		Set(ts.Temp)
	}
	saveStatusFile(strconv.Itoa(temp.Value))
}
