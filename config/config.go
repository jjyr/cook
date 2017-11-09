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
	return
}
