package cmdproxy

import (
	"golang.org/x/crypto/ssh"
	"io"
	"github.com/jjyr/cook/common"
	"io/ioutil"
	"net"
	"strings"
	"fmt"
	"os/user"
)

type SSHProxyClient struct {
	Stdin   io.WriteCloser
	Stdout  io.Reader
	Stderr  io.Reader
	client  *ssh.Client
	session *ssh.Session
}

var _ ProxyClient = &SSHProxyClient{}

func setServerDefaultValues(server common.Server) (common.Server) {
	if server.PrivateKeyFile == "" {
		server.PrivateKeyFile = "~/.ssh/id_rsa"
	}
	if server.User == "" {
		u, err := user.Current()
		if err != nil {
			panic(fmt.Errorf("can't get current os user %s", err))
		}
		server.User = u.Username
	}
	if server.Port == "" {
		server.Port = "22"
	}
	return server
}

func getServerAuthMethods(server common.Server) (authMethods []ssh.AuthMethod, err error) {
	if server.PassWord != "" {
		authMethods = []ssh.AuthMethod{
			ssh.Password(server.PassWord),
		}
		return
	}

	privateKey, err := ioutil.ReadFile(server.PrivateKeyFile)
	if err != nil {
		return
	}
	signer, err := ssh.ParsePrivateKey([]byte(privateKey))
	if err != nil {
		return
	}
	authMethods = []ssh.AuthMethod{
		ssh.PublicKeys(signer),
	}
	return
}

func newSSHClient(server common.Server) (client *ssh.Client, err error) {
	server = setServerDefaultValues(server)
	authMethods, err := getServerAuthMethods(server)
	if err != nil {
		return
	}

	clientConfig := &ssh.ClientConfig{
		User: server.User,
		Auth: authMethods,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil //ignore
		},
	}

	client, err = ssh.Dial("tcp", fmt.Sprintf("%s:%s", server.Host, server.Port), clientConfig)
	return
}

func NewSSHProxyClient(server common.Server) (dockerClient *SSHProxyClient, err error) {
	dockerClient = &SSHProxyClient{}

	client, err := newSSHClient(server)
	if err != nil {
		return
	}
	dockerClient.client = client
	session, err := client.NewSession()
	if err != nil {
		return
	}
	dockerClient.session = session
	if dockerClient.Stdout, err = session.StdoutPipe(); err != nil {
		return
	}

	if dockerClient.Stderr, err = session.StderrPipe(); err != nil {
		return
	}

	return
}

func (client *SSHProxyClient) Start(name string, args ... string) error {
	cmd := fmt.Sprintf("%s %s", name, strings.Join(args, " "))
	return client.session.Start(cmd)
}

func (client *SSHProxyClient) Wait() error {
	return client.session.Wait()
}

func (client *SSHProxyClient) StdinPipe() (w io.WriteCloser) {
	if client.Stdin != nil {
		w = client.Stdin
		return
	}
	w, err := client.session.StdinPipe()
	if err != nil {
		panic(err)
	}
	client.Stdin = w
	return
}

func (client *SSHProxyClient) StdoutPipe() (io.Reader) {
	if client.Stdout != nil {
		return client.Stdout
	}
	var err error
	client.Stdout, err = client.session.StdoutPipe()
	if err != nil {
		panic(err)
	}
	return client.Stdout
}

func (client *SSHProxyClient) StderrPipe() (r io.Reader) {
	if client.Stderr != nil {
		r = client.Stderr
		return
	}
	r, err := client.session.StderrPipe()
	if err != nil {
		panic(err)
	}
	client.Stderr = r
	return
}

func (client *SSHProxyClient) Run(name string, args ... string) error {
	cmd := fmt.Sprintf("%s %s", name, strings.Join(args, " "))
	return client.session.Run(cmd)
}

func (client *SSHProxyClient) Close() (err error) {
	if client.session != nil {
		if err = client.session.Close(); err != nil {
			return
		}
	}
	if client.client != nil {
		err = client.client.Close()
	}
	return
}
