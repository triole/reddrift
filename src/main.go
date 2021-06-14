package main

func main() {
	parseArgs()

	if CLI.Presets == true {
		listPresets()
	} else {
		setTemp()
	}
}
