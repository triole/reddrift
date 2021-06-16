package main

import (
	"testing"
)

func TestAutoCalcTemp(t *testing.T) {
	for i := 1; i <= 200; i = i + 5 {
		ts.SunAltitude = float64(float64(i-100) / 100)
		_ = autoCalcTemp(ts, 2500, 6500)
		// fmt.Printf("%+.3f -- %d\n", ts.SunAltitude, temp)
	}
}
