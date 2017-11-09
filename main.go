package main

import (
	"github.com/urfave/cli"
	"github.com/jjyr/cook/cmd"
	"os"
	log "github.com/sirupsen/logrus"
)

func main() {
	app := cmd.InitApp()

	app.Before = func(context *cli.Context) error {
		log.SetOutput(os.Stdout)
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
