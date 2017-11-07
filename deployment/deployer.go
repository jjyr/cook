package deployment

import (
	"github.com/jjyr/cook/common"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"net"
	"errors"
	"io"
)

type Deployer struct {
	Server common.Server
}

func NewDeployer(server common.Server) *Deployer {
	return &Deployer{Server: server}
}

const privateKeyFile = `/Users/jiangjinyang/VMs/ubuntu/.vagrant/machines/default/virtualbox/private_key`

func (d *Deployer) Prepare(image common.Image) (err error) {
	//cmd := exec.Command("docker-compose", args...)
	//err = cmd.Run()
	// get from docker compose
	privateKey, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		panic(err)
	}
	signer, err := ssh.ParsePrivateKey([]byte(privateKey))
	if err != nil {
		panic("Failed to parse private key: " + err.Error())
	}
	clientConfig := &ssh.ClientConfig{
		User: "vagrant",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil //ignore
		},
	}
	client, err := ssh.Dial("tcp", "127.0.0.1:2222", clientConfig)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}
	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()
	w, _ := session.StdinPipe()
	defer w.Close()
	r, _ := session.StdoutPipe()
	err = session.Start("/bin/cat")
	if err != nil {
		panic("Failed to run: " + err.Error())
	}

	fmt.Print("writing something\n")
	fmt.Fprintln(w, "Hello world ok?")
	fmt.Fprint(w, io.EOF)
	w.Close()

	fmt.Print("writed, wating..\n")

	err = session.Wait()
	if err != nil {
		panic("Failed to wait?: " + err.Error())
	}

	out, err := ioutil.ReadAll(r)
	if err != nil {
		panic("Failed to read?: " + err.Error())
	}
	fmt.Print("output is:", string(out))
	err = errors.New("fuck done")
	return
}

func (d *Deployer) Deploy() (err error) {
	// docker run docker-compose < compose-file
	return
}
