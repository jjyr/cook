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
	app.Usage = "Simple deployment workflow for docker/compose, save poor man's life"
	app.Description = `cook provide a simple workflow for docker-compose user:
	1) build images on local-machine
	2) push images to target servers
	3) run 'docker-compose up -d' on target servers
	4) check services status, make sure they are 'up'

	you need cook.yml file to configure server and docker-compose

	type 'cook' to run default workflow
	type 'cook config --sample' to see how to write a cook.yml file
	`
	app.Action = commands.Main
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Config file",
			Value: "./cook.yml",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "build",
			Usage:  "Build docker images",
			Action: commands.Build,
		}, {
			Name:   "prepare",
			Usage:  "Push docker images to target servers",
			Action: commands.Prepare,
		}, {
			Name:   "deploy",
			Usage:  "Run 'docker-compose up' on target servers",
			Action: commands.Deploy,
		}, {
			Name:  "config",
			Usage: "Display full configuration with config file",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "sample, s"},
			},
			Action: commands.Config,
		},
	}
	return
}
