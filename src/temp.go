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
		if err != nil {
			return t, fmt.Errorf("can not parse int")
		}
		if i < presets["minimal"] || i > presets["maximal"] {
			return t, fmt.Errorf(
				"temp out of range, value needs to be between %v and %v",
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

func setTemp() {
	temp := Temp{}
	temp, err := temp.Set(CLI.Temp)
	if err != nil {
		fmt.Printf("%q\n", err)
	} else {
		fmt.Printf("Set color temperatur to %q %v\n", temp.Name, temp.Value)
		Set(temp.Value)
	}
}
