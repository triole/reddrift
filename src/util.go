package main

import (
	"sort"
	"strings"
)

func mapKeysToString(m map[string]int) string {
	hackMap, hackKeys := sortMapHack(m)
	keys := make([]string, 0, len(m))
	for _, k := range hackKeys {
		keys = append(keys, hackMap[k])
	}
	return "[" + strings.Join(keys, ", ") + "]"
}

func sortMapHack(m map[string]int) (hackMap map[int]string, hackKeys []int) {
	hackMap = make(map[int]string)
	for key, val := range m {
		hackMap[val] = key
		hackKeys = append(hackKeys, val)
	}
	sort.Ints(hackKeys)
	return hackMap, hackKeys
}
