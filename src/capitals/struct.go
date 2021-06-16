package capitals

// Capitals holds the capitals class
type Capitals struct {
	Capitals tCapitals
}

type tLocation struct {
	Capital string
	Country string
	Coords  tCoords
}

type tCoords struct {
	Lat float64
	Lon float64
}

type tCapitals []tLocation

func (tcap tCapitals) Len() int {
	return len(tcap)
}

func (tcap tCapitals) Less(i, j int) bool {
	return tcap[i].Capital < tcap[j].Capital
}

func (tcap tCapitals) Swap(i, j int) {
	tcap[i], tcap[j] = tcap[j], tcap[i]
}
