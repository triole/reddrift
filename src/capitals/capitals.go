package capitals

import (
	_ "embed"
	"encoding/json"
	"strings"
)

//go:embed capitals.json
var capitalsEmbed string

// Self holds the capitals class
type Self struct {
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

// Init does what it says
func Init() (s Self) {
	s.Capitals = readJSON(capitalsEmbed)
	return s
}

func readJSON(content string) (capitals tCapitals) {
	err := json.Unmarshal([]byte(content), &capitals)
	if err != nil {
		panic(err)
	}
	return
}

func (self Self) GetLocation(s string) (loc tLocation) {
	for _, cap := range self.Capitals {
		if strings.ToLower(s) == strings.ToLower(cap.Capital) {
			loc = cap
			break
		}
	}
	return
}
