package deployment

import (
	"github.com/jjyr/cook/common"
	"fmt"
	"io/ioutil"
	"io"
	"github.com/sirupsen/logrus"
	"github.com/jjyr/cook/cmdproxy"
)

type Deployer struct {
	Server common.Server
}

func NewDeployer(server common.Server) *Deployer {
	return &Deployer{Server: server}
}

func (d *Deployer) Prepare(images ... common.Image) (err error) {
	var remoteDocker, localDocker cmdproxy.ProxyClient
	remoteDocker, err = cmdproxy.NewSSHProxyClient(d.Server)
	if err != nil {
		return err
	}
	defer remoteDocker.Close()
	remoteStdin := remoteDocker.StdinPipe()
	defer remoteStdin.Close()
	localDocker, _ = cmdproxy.NewLocalProxyClient()
	defer localDocker.Close()
	// remote: docker load
	err = remoteDocker.Start("docker load")
	if err != nil {
		panic("Fail to run remote command: " + err.Error())
	}
	// local: docker save
	err = localDocker.Start("docker", append([]string{"save"}, images...)...)
	if err != nil {
		panic(err)
	}

	// copy local images to remote
	size, err := io.Copy(remoteStdin, localDocker.StdoutPipe())
	if err != nil {
		panic(err)
	}

	logrus.Debugf("written %d to remote\n", size)
	remoteStdin.Close()

	localError, err := ioutil.ReadAll(localDocker.StderrPipe())
	if len(localError) > 0 {
		panic(fmt.Errorf("docker save error: %s\n", string(localError)))
	}

	err = remoteDocker.Wait()
	if err != nil {
		panic("Failed to wait remote docker: " + err.Error())
	}
	err = localDocker.Wait()
	if err != nil {
		panic("Failed to wait local docker: " + err.Error())
	}

	return
}

func (d *Deployer) Deploy() (err error) {
	// docker run docker-compose < compose-file
	return
}
