package controller

import (
	"github.com/jjyr/cook/config"
	"github.com/jjyr/cook/deployment"
	"github.com/jjyr/cook/log"
	"github.com/jjyr/cook/common"
	"github.com/jjyr/cook/backend"
	"os"
)

type Controller struct {
	config.Config
	Logger log.Logger
}

func NewController(c config.Config) *Controller {
	return &Controller{Config: c}
}

func (c *Controller) Build() (err error) {
	// run docker-compose build

	for _, deployDesc := range c.Config.Deploy {
		d := backend.GetBackend(deployDesc)
		c.Logger.Infof("Run build %s", d.Path)
		cmd := d.BuildCmd()
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err = cmd.Run(); err != nil {
			c.Logger.Fatal(err)
		}
		c.Logger.Infoln("Build complete")
	}
	return
}

func (c *Controller) Prepare() (err error) {
	// get images from docker compose
	// push images to remote target

	for _, server := range c.Config.Target {
		images := make([]common.Image, 0)
		for _, deployDesc := range c.Config.Deploy {
			deployImages, err := backend.GetBackend(deployDesc).Images()
			if err != nil {
				c.Logger.Fatal(err)
			}
			images = append(images, deployImages...)
		}

		deployer := deployment.NewDeployer(server)
		deployer.Logger = c.Logger
		c.Logger.Infof("Prepare push images %s to server %s", images, server.String())
		if err = deployer.Prepare(images...); err != nil {
			return
		}
		c.Logger.Infoln("Push complete")
	}
	return
}

func (c *Controller) Deploy() (err error) {
	// get image from docker compose
	// run deploy-images on remote machines
	// done
	for _, server := range c.Config.Target {
		deployer := deployment.NewDeployer(server)
		deployer.Logger = c.Logger

		for _, deployDesc := range c.Config.Deploy {
			c.Logger.Infof("Deploy %s to %s", deployDesc.Path, server.String())
			err = deployer.Deploy(deployDesc)
			if err != nil {
				c.Logger.Fatal(err)
			}
		}
		c.Logger.Infoln("Done")
	}
	return
}
