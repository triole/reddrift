package main

// TempSet bundles all relevant data, basically it is required for logging
type TempSet struct {
	Capital     string
	Country     string
	Lat         float64
	Lon         float64
	SunAltitude float64
	SunAzimuth  float64
	Temp        int
	TempName    string
	LastTemp    int
}
