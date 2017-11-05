package deployment

import "github.com/jjyr/cook/config/cook"

type Deployment struct {
	cook.CookConfig
}

func NewDeployment(c cook.CookConfig) *Deployment {
	return &Deployment{CookConfig: c}
}

func (d *Deployment) Prepare() (err error) {
	return
}

func (d *Deployment) Deploy() (err error) {
	return
}
