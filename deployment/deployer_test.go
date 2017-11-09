package deployment

import (
	"testing"
	"github.com/jjyr/cook/common"
	"fmt"
	"net"
)

var testServer = common.Server{
	PrivateKeyFile: "/Users/jiangjinyang/VMs/ubuntu/.vagrant/machines/default/virtualbox/private_key",
	Host:           "127.0.0.1",
	Port:           "2222",
	User:           "root",
}

func canPingServer(server common.Server) bool {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", server.Host, server.Port))
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func TestDeployer_Prepare(t *testing.T) {
	if !canPingServer(testServer) {
		t.Skip("can not connect to test server %+v", testServer)
		return
	}
	deploy := NewDeployer(testServer)
	if err := deploy.Prepare("mysql:latest"); err != nil {
		t.Error(err)
	}
}

func TestDeployer_Deploy(t *testing.T) {
	if !canPingServer(testServer) {
		t.Skip("can not connect to test server %+v", testServer)
		return
	}
	deploy := NewDeployer(testServer)
	desc := common.DeployDesc{
		Path: "/Users/jiangjinyang/workspace/go/src/github.com/jjyr/cook/test/docker-compose.yml",
	}
	if err := deploy.Deploy(desc); err != nil {
		t.Error(err)
	}
}
