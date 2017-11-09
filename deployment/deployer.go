package deployment

import (
	"github.com/jjyr/cook/common"
	"fmt"
	"io/ioutil"
	"io"
	"github.com/jjyr/cook/cmdproxy"
	"os"
	"path"
	"github.com/jjyr/cook/log"
)

type Deployer struct {
	Server common.Server
	Logger log.Logger
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

	d.Logger.Infof("written %d to remote\n", size)
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

func (d *Deployer) Deploy(desc common.DeployDesc) (err error) {
	// docker run docker-compose < compose-file
	// docker-compose mkdir docker-compose
	projectName := desc.ProjectName
	if projectName == "" {
		projectName = path.Base(path.Dir(desc.Path))
	}
	file, err := os.Open(desc.Path)
	if err != nil {
		return
	}
	defer file.Close()

	var remoteDocker cmdproxy.ProxyClient
	remoteDocker, err = cmdproxy.NewSSHProxyClient(d.Server)
	if err != nil {
		return
	}

	remoteStdin := remoteDocker.StdinPipe()
	defer remoteStdin.Close()

	err = remoteDocker.Start("docker", "run", "-i", "--rm", "-v", "/var/run/docker.sock:/var/run/docker.sock",
		"jjy0/docker-compose-for-cook:latest", "/bin/bash", "-c",
		fmt.Sprintf("'cat - > docker-compose.yml && docker-compose --project-name %s up -d --no-build'", projectName))
	if err != nil {
		panic(fmt.Errorf("docker-compose execute failed: %s\n", err))
	}
	if _, err = io.Copy(remoteStdin, file); err != nil {
		panic(fmt.Errorf("write to remote error: %s", err))
	}
	if err = remoteStdin.Close(); err != nil {
		panic(err)
	}
	err = remoteDocker.Wait()

	out, _ := ioutil.ReadAll(remoteDocker.StdoutPipe())
	errOut, _ := ioutil.ReadAll(remoteDocker.StderrPipe())
	d.Logger.Infof("remote server output:%s error output:%s\n", string(out), string(errOut))

	if err != nil {
		panic("Failed to wait remote docker: " + err.Error())
	}
	return
}
