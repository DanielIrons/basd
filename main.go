package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

const (
	buildDate = "unknown"
	buildHash = "unknown"
)

func main() {
	compiled, _ := time.Parse("", buildDate)
	app := cli.App{
		Name:                   "basd",
		UseShortOptionHandling: true,
		EnableBashCompletion:   true,
		Suggest:                true,
		Compiled:               compiled,
		Version:                buildHash,
		Authors: []*cli.Author{
			{
				Name: "Daniel Irons",
			},
		},
		Usage: "basic *** service discovery",
		Commands: []*cli.Command{
			{
				Name:    "server",
				Aliases: []string{"s"},
				Usage:   "start the basd server",
				Action: func(cCtx *cli.Context) error {
					log.Println("running server")
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
