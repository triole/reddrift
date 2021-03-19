package main

import (
	"fmt"
	"os"
)

func main() {
	argparse()

	if *argsPresets == true {
		fmt.Println("\nAvailable presets")
		hackedMap, hackedKeys := sortMapHack(presets)
		for _, k := range hackedKeys {
			fmt.Printf("%v\t%s\n", k, hackedMap[k])
		}
		fmt.Printf("\n")
		os.Exit(0)
	}

	temp := Temp{}
	temp, err := temp.Set(*argsTemp)

	if err != nil {
		fmt.Printf("%q\n", err)
	} else {
		fmt.Printf("Set color temperatur to %q %v\n", temp.Name, temp.Value)
		Set(temp.Value)
	}
}
