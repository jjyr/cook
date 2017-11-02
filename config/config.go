package config

import (
	"github.com/jjyr/cook/config/cook"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	cook.CookConfig
}

func LoadConfig(path string) (c Config, err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(data, &(c.CookConfig))
	return
}
