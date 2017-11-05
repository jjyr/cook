package config

import (
	"github.com/jjyr/cook/config/cook"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	cook.CookConfig
	DeployName string
}

func LoadConfig(path string, deployName string) (c Config, err error) {
	c.DeployName = deployName
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(data, &(c.CookConfig))
	return
}
