package main

import (
	"fmt"
	"os"

	"github.com/mradile/rssfeeder/pkg/client/cmd"

	"github.com/urfave/cli"
)

var version = "dev-snapshot"

var debug bool
var verbose bool

func main() {

	cliApp := cli.NewApp()
	cliApp.Name = "rssfeeder client"
	cliApp.Usage = ""
	cliApp.Version = version

	cliApp.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "prints debug level log messages",
			EnvVar:      "LOG_DEBUG",
			Destination: &debug,
		},
		cli.BoolFlag{
			Name:        "verbose",
			Usage:       "prints info level log messages",
			EnvVar:      "LOG_VERBOSE",
			Destination: &verbose,
		},
		cli.StringFlag{
			Name:   "cfg",
			Usage:  "config `directory`",
			EnvVar: "CONFIG_DIRECTORY",
			Value:  "",
		},
	}

	cmds := make([]cli.Command, 0)
	cmds = append(cmds, cmd.Client.Login())
	cmds = append(cmds, cmd.Client.AddEntry())
	cmds = append(cmds, cmd.Client.ListFeeds())
	cliApp.Commands = cmds

	err := cliApp.Run(os.Args)
	if err != nil {
		panic(fmt.Sprintf("could not initialize app: %s", err))
	}
}
