package main

import "time"

// TempSet bundles all relevant data, basically it is required for logging
type tempSet struct {
	Date          time.Time
	Capital       string
	Country       string
	Lat           float64
	Lon           float64
	CurveModifier float64
	CurveModList  []float64
	SunAltitude   float64
	SunAzimuth    float64
	Temp          int
	TempName      string
	LastTemp      int
	ValidPreset   bool
	ValidTempInt  bool
}
