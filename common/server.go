package common

import "fmt"
import "os/user"

type Server struct {
	PrivateKeyFile string `yaml:"private_key_file"`
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	User           string `yaml:"user"`
	PassWord       string `yaml:"password"`
}

var defaultServer Server

func SetServerDefault(s *Server) {
	if s.PrivateKeyFile == "" {
		s.PrivateKeyFile = defaultServer.PrivateKeyFile
	}
	if s.User == "" {
		s.User = defaultServer.User
	}
	if s.Port == "" {
		s.Port = defaultServer.Port
	}
}

func init() {
	u, err := user.Current()
	if err != nil {
		panic(fmt.Errorf("can't get current os user %s", err))
	}

	defaultServer = Server{
		PrivateKeyFile: "~/.ssh/id_rsa",
		User:           u.Username,
		Port:           "22",
	}
}
