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
