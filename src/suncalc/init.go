package suncalc

import (
	"time"

	"github.com/sixdouglas/suncalc"
)

// SunData holds the result of the sun position calculation
type SunData struct {
	Time     time.Time
	Lat      float64
	Lon      float64
	Altitude float64
	Azimuth  float64
}

// Init does exactly what it says
func Init(lat, lon float64) (sd SunData) {
	sd.Time = time.Now()
	sd.Lat = lat
	sd.Lon = lon
	r := suncalc.GetPosition(sd.Time, sd.Lat, sd.Lon)
	sd.Altitude = r.Altitude
	sd.Azimuth = r.Azimuth
	return
}
