package main

import (
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	BUILDTAGS      string
	appName        = "sct"
	appMainVersion = "0.1"
	appDescription = "sct"

	app      = kingpin.New(appName, appDescription)
	argsTemp = app.Arg(
		"temp", "temperature to set as int, or by preset "+mapKeysToString(presets),
	).Default("6500").String()
	argsPresets = app.Flag("presets", "list available presets and their value").Short('p').Bool()
)

func argparse() {
	env := tEnv{
		Name:        appName,
		MainVersion: appMainVersion,
		Description: appDescription,
	}
	app.Version(makeInfoString(env, parseBuildtags(BUILDTAGS)))
	app.VersionFlag.Short('V')
	app.HelpFlag.Short('h')

	kingpin.MustParse(app.Parse(os.Args[1:]))
}
