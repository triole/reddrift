package main

import (
	"fmt"
	"strconv"
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

func setTemp(tempInt string) {
	temp := Temp{}
	temp, err := temp.Set(tempInt)
	if err != nil {
		fmt.Printf("%q\n", err)
	} else {
		if temp.Name == "manual" {
			lg.Log("Set colour temp to %d\n", temp.Value)
		} else {
			lg.Log("Preset "+temp.Name+", set colour temp %d\n", temp.Value)
		}
		Set(temp.Value)
	}
}
