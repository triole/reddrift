package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/alecthomas/kong"
)

var (
	BUILDTAGS      string
	appName        = "sct"
	appDescription = "set color temperature is a very simplt replacement for redshift"
	appMainversion = "0.1"
)

var CLI struct {
	Temp        string `help:"color temp or preset to set" arg default:6500k`
	Presets     bool   `help:"list available presets" short:p`
	VersionFlag bool   `help:"display version" short:V`
}

func parseArgs() {
	ctx := kong.Parse(&CLI,
		kong.Name(appName),
		kong.Description(appDescription),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}),
	)
	_ = ctx.Run()

	if CLI.VersionFlag == true {
		printBuildTags(BUILDTAGS)
		os.Exit(0)
	}
}

func printBuildTags(buildtags string) {
	regexp, _ := regexp.Compile(`({|}|,)`)
	s := regexp.ReplaceAllString(buildtags, "\n")
	s = strings.Replace(s, "_subversion: ", "Version: "+appMainversion+".", -1)
	fmt.Printf("%s\n", s)
}

func alnum(s string) string {
	s = strings.ToLower(s)
	re := regexp.MustCompile("[^a-z0-9_-]")
	return re.ReplaceAllString(s, "-")
}
