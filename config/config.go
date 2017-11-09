package config

import (
	"github.com/jjyr/cook/common"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Target []common.Server     `yaml:"target"`
	Deploy []common.DeployDesc `yaml:"deploy"`
}

func LoadConfig(path string) (c Config, err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return
	}
	SetConfigDefault(&c)
	return
}

func SetConfigDefault(c *Config) {
	for i, _ := range c.Target {
		common.SetServerDefault(&(c.Target[i]))
	}
	for i, _ := range c.Deploy {
		common.SetDeployDescDefault(&(c.Deploy[i]))
	}
}
