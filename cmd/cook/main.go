package main

import (
	"github.com/urfave/cli"
	"github.com/jjyr/cook/logger"
	"github.com/jjyr/cook/cmd"
	"os"
)

func main() {
	app := cmd.InitApp()

	log := logger.New()

	app.Before = func(context *cli.Context) error {
		log.Out = os.Stdout
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
