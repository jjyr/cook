package config

import (
	"github.com/jjyr/cook/common"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"github.com/jjyr/cook/backend"
)

type Config struct {
	Target     []common.Server      `yaml:"target"`
	Deploy     []backend.DeployDesc `yaml:"deploy"`
	DeployName string               `yaml:"name"`
}

func LoadConfig(path string, deployName string) (c Config, err error) {
	c.DeployName = deployName
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(data, &c)
	return
}
