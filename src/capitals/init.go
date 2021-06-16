package capitals

import (
	_ "embed"
	"sort"
)

//go:embed capitals.json
var capitalsEmbed string

// Init does what it says
func Init() (cap Capitals) {
	cap.Capitals = readJSON(capitalsEmbed)
	sort.Sort(cap.Capitals)
	return cap
}
