package controller

import (
	"github.com/jjyr/cook/config"
	"github.com/jjyr/cook/builder"
	"github.com/jjyr/cook/deployment"
)

type Controller struct {
	config.Config
}

func NewController(c config.Config) *Controller {
	return &Controller{Config: c}
}

func (c *Controller) Build() (err error) {
	b := builder.NewBuilder(c.CookConfig.Build.Dockerfiles[0])
	err = b.Build()
	// build 描述 Dockerfile
	return
}

func (c *Controller) Prepare() (err error) {
	d := deployment.NewDeployment(c.CookConfig)
	err = d.Prepare()
	return
}

func (c *Controller) Deploy() (err error) {
	d := deployment.NewDeployment(c.CookConfig)
	err = d.Deploy()
	return
}
