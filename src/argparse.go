package main

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/alecthomas/kong"
)

var (
	BUILDTAGS      string
	appName        = "sct"
	appDescription = "set colour temperature is a very simplt replacement for redshift"
	appMainversion = "0.1"
)

var CLI struct {
	Temp         string `help:"colour temp or preset to set" arg default:6500`
	Presets      bool   `help:"list available presets" short:p`
	Auto         bool   `help:"auto set temp for default location" short:a`
	Min          int    `help:"auto mode minimum" short:m default:2500`
	Max          int    `help:"auto mode maximum" short:m default:6500`
	Location     string `help:"custom location, currently supported capitals (i.e. london, paris)" short:c default:Berlin`
	LogFile      string `help:"log file" short:l default:${logfile}`
	StatusFile   string `help:"status file" short:s default:${statusfile}`
	Repeat       bool   `help:"keep running and auto adjust continuously" short:r`
	TickInterval int    `help:"tick interval when repeat enabled, check every x seconds" short:t default:10`
	VersionFlag  bool   `help:"display version" short:V`
}

func parseArgs() {
	an := alnum(appName)
	fol := path.Join(os.TempDir(), an)
	os.Mkdir(fol, 0755)
	ctx := kong.Parse(&CLI,
		kong.Name(appName),
		kong.Description(appDescription),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}),
		kong.Vars{
			"logfile":    path.Join(fol, an+".log"),
			"statusfile": path.Join(fol, an+"_status.txt"),
		},
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
