package cmdproxy

import (
	"io"
	"os/exec"
	"fmt"
)

type LocalProxyClient struct {
	Stdin  io.WriteCloser
	Stdout io.ReadCloser
	Stderr io.ReadCloser
	cmd    *exec.Cmd
}

var _ ProxyClient = &LocalProxyClient{}

func NewLocalProxyClient() (dockerClient *LocalProxyClient, err error) {
	dockerClient = &LocalProxyClient{}
	return
}

func (client *LocalProxyClient) newCmd(name string, args ... string) {
	if client.cmd != nil {
		panic(fmt.Errorf("previous command not finished:%+v", client.cmd))
	}
	client.cmd = exec.Command("docker", args...)
	var err error
	client.Stderr, err = client.cmd.StderrPipe()
	if err != nil {
		panic(err)
	}
	client.Stdout, err = client.cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
}

func (client *LocalProxyClient) Start(name string, args ... string) error {
	client.newCmd(name, args...)
	return client.cmd.Start()
}

func (client *LocalProxyClient) Wait() error {
	return client.cmd.Wait()
}

func (client *LocalProxyClient) StdinPipe() (w io.WriteCloser) {
	panic("not implemented")
	return
}

func (client *LocalProxyClient) StdoutPipe() (r io.Reader) {
	if client.Stdout != nil {
		r = client.Stdout
		return
	}
	var err error
	client.Stdout, err = client.cmd.StdoutPipe()
	r = client.Stdout
	if err != nil {
		panic(err)
	}
	return
}

func (client *LocalProxyClient) StderrPipe() (r io.Reader) {
	if client.Stderr != nil {
		r = client.Stderr
		return
	}
	var err error
	client.Stderr, err = client.cmd.StderrPipe()
	r = client.Stderr
	if err != nil {
		panic(err)
	}
	return
}

func (client *LocalProxyClient) Run(name string, args ... string) error {
	client.newCmd(name, args...)
	return client.cmd.Run()
}

func (client *LocalProxyClient) Close() (err error) {
	if client.cmd == nil {
		return
	}
	if client.Stdout != nil {
		err = client.Stdout.Close()
	}
	if client.Stderr != nil {
		err = client.Stderr.Close()
	}
	client.cmd = nil
	return
}
