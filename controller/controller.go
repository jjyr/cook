package controller

import (
	"github.com/jjyr/cook/config"
	"github.com/jjyr/cook/deployment"
	log "github.com/sirupsen/logrus"
	"github.com/jjyr/cook/common"
)

type Controller struct {
	config.Config
	Logger *log.Logger
}

func NewController(c config.Config) *Controller {
	return &Controller{Config: c}
}

func (c *Controller) Build() (err error) {
	// run docker-compose build

	for _, deployDesc := range c.Config.Deploy {
		d := deployDesc.GetBackend()
		c.Logger.Infof("Run build from %s\n", d.Path)
		if err = d.Build(); err != nil {
			c.Logger.Fatal(err)
		}
		c.Logger.Infoln("Done")
	}
	return
}

func (c *Controller) Prepare() (err error) {
	// get images from docker compose
	// push images to remote target

	for _, server := range c.Config.Target {
		c.Logger.Infof("Prepare %s images\n", server)

		images := make([]common.Image, 0)
		for _, deployDesc := range c.Config.Deploy {
			deployImages, err := deployDesc.GetBackend().Images()
			if err != nil {
				c.Logger.Fatal(err)
			}
			images = append(images, deployImages...)
		}

		deployer := deployment.NewDeployer(server)

		for _, image := range images {
			c.Logger.Infof("Push %s\n", image)
			if err = deployer.Prepare(image); err != nil {
				return
			}
		}
		c.Logger.Infoln("Done")
	}
	return
}

func (c *Controller) Deploy() (err error) {
	// get image from docker compose
	// run deploy-images on remote machines
	// done
	d := deployment.NewDeployer(c.Config)
	err = d.Deploy()
	return
}
