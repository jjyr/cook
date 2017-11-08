package common

type Server struct {
	PrivateKeyFile string `yaml:"private_key_file"`
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	User           string `yaml:"user"`
	PassWord       string `yaml:"password"`
}
