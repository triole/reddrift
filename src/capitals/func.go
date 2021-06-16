package capitals

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (cap Capitals) GetLocation(s string) (loc tLocation) {
	for _, cap := range cap.Capitals {
		if strings.ToLower(s) == strings.ToLower(cap.Capital) {
			loc = cap
			break
		}
	}
	return
}

func (cap Capitals) ListLocations() {
	for _, cap := range cap.Capitals {
		fmt.Printf("%+v\n", cap)
	}
}

func readJSON(content string) (capitals tCapitals) {
	err := json.Unmarshal([]byte(content), &capitals)
	if err != nil {
		panic(err)
	}
	return
}
