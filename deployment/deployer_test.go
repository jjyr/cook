package deployment

import (
	"testing"
	"github.com/jjyr/cook/common"
)

func TestDeployer_Prepare(t *testing.T) {
	deploy := NewDeployer(common.Server{
		PrivateKeyFile: "/Users/jiangjinyang/VMs/ubuntu/.vagrant/machines/default/virtualbox/private_key",
		Host:           "127.0.0.1",
		Port:           "2222",
		User:           "root",
	})
	if err := deploy.Prepare("mysql:latest"); err != nil {
		t.Error(err)
	}
}

func TestDeployer_Deploy(t *testing.T) {
	deploy := NewDeployer(common.Server{
		PrivateKeyFile: "/Users/jiangjinyang/VMs/ubuntu/.vagrant/machines/default/virtualbox/private_key",
		Host:           "127.0.0.1",
		Port:           "2222",
		User:           "root",
	})
	desc := common.DeployDesc{
		Path: "/Users/jiangjinyang/workspace/go/src/github.com/jjyr/cook/test/docker-compose.yml",
	}
	if err := deploy.Deploy(desc); err != nil {
		t.Error(err)
	}
}
