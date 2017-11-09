package main

import (
	"github.com/urfave/cli"
	"github.com/jjyr/cook/cmd"
	"os"
	"github.com/sirupsen/logrus"
)

func main() {
	app := cmd.InitApp()

	app.Before = func(context *cli.Context) error {
		logrus.SetOutput(os.Stdout)
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
