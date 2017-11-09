package config

import (
	"testing"
	"github.com/jjyr/cook/common"
)

func TestSetConfigDefault(t *testing.T) {
	c := Config{
		Target: []common.Server{{}},
		Deploy: []common.DeployDesc{{}},
	}
	SetConfigDefault(&c)
	if c.Target[0] == (common.Server{}) || c.Deploy[0] == (common.DeployDesc{}) {
		t.Error("SetConfigDefault not work")
	}
}
