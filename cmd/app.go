package cmd

import (
	"github.com/urfave/cli"
	"github.com/jjyr/cook/config"
	"github.com/jjyr/cook/cmd/commands"
)

func InitApp() (app *cli.App) {
	app = cli.NewApp()
	app.Name = "cook"
	app.Version = config.Version
	app.Usage = "Simple deployment workflow with docker, save poor man's life"
	app.Action = commands.Main
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from config file, default is `./cook.yml`",
			Value: "./cook.yml",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:        "build",
			Usage:       "cook build [dockerfile]",
			Description: "Build docker images",
			Action:      commands.Build,
		}, {
			Name:        "prepare",
			Usage:       "cook prepare image [servers...]",
			Description: "Push docker image to target servers",
			Action:      commands.Prepare,
		}, {
			Name:        "deploy",
			Usage:       "cook deploy deploy-name [servers...]",
			Description: "Apply deployment to target servers",
			Action:      commands.Deploy,
		}, {
			Name:        "config",
			Usage:       "cook config",
			Description: "Display full configuration with config file",
			Action:      commands.Config,
		},
	}
	return
}
