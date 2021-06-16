package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func absPath(str string) string {
	p, err := filepath.Abs(str)
	lg.LogIfErrFatal(err, "Invalid file path %q", str)
	return p
}

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

func saveStatusFile(val string) {
	data := []byte(val)
	err := ioutil.WriteFile(CLI.StatusFile, data, 0644)
	lg.LogIfErr(err, "Can not save status file %q", CLI.StatusFile)
}

func readStatusFile() (temp int) {
	var content []byte
	var err error
	content, err = ioutil.ReadFile(CLI.StatusFile)
	lg.LogIfErr(err, "Can not read status file %q", CLI.StatusFile)
	temp, _ = strconv.Atoi(string(content))
	return
}

func listPresets() {
	fmt.Println("\nAvailable presets")
	hackedMap, hackedKeys := sortMapHack(presets)
	for _, k := range hackedKeys {
		fmt.Printf("%v\t%s\n", k, hackedMap[k])
	}
	fmt.Printf("\n")
}
