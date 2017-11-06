package deployment

import (
	"github.com/jjyr/cook/config"
)

type Deployment struct {
	config.Config
}

func NewDeployment(c config.Config) *Deployment {
	return &Deployment{Config: c}
}

func (d *Deployment) Prepare() (err error) {
	return
}

func (d *Deployment) Deploy() (err error) {
	return
}
