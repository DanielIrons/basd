package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ironsdan/basd/config"
	"github.com/urfave/cli/v2"
)

var (
	buildDate = "unknown"
	buildHash = "unknown"
)

func main() {
	compiled, _ := time.Parse("", buildDate)

	cli.VersionPrinter = func(cCtx *cli.Context) {
		fmt.Printf("%s built %s\n", cCtx.App.Version, cCtx.App.Compiled)
	}
	app := cli.App{
		Name:                   "basd",
		UseShortOptionHandling: true,
		EnableBashCompletion:   true,
		Suggest:                true,
		Compiled:               compiled,
		Version:                buildHash,
		Authors: []*cli.Author{
			{
				Name:  "Daniel Irons",
				Email: "ironsdan@pm.me",
			},
		},
		Usage: "basic service discovery",
		Commands: []*cli.Command{
			{
				Name:    "agent",
				Aliases: []string{"a"},
				Usage:   "start the basd agent",
				Action: func(cCtx *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "show version",
				Action: func(cCtx *cli.Context) error {
					cli.VersionPrinter(cCtx)
					return nil
				},
			},
		},
		Before: func(cCtx *cli.Context) error {
			c, err := config.LoadConfig(cCtx.String("config"))
			if err != nil {
				return err
			}
			cCtx.App.Metadata["config"] = c
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "config", Aliases: []string{"c"}, Usage: "path to config file", Value: "config.json"},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
