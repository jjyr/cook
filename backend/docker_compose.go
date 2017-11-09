package backend

import (
	"os/exec"
	"github.com/jjyr/cook/common"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
)

type DockerCompose struct {
	Path    string
	WorkDir string
	Compose
}

type Service struct {
	Image string `yaml:"image"`
	Build string `yaml:"build"`
}

type Compose struct {
	Version  string             `yaml:"version"`
	Services map[string]Service `yaml:"services"`
}

func NewDockerCompose() (*DockerCompose) {
	return &DockerCompose{Path: "./docker-compose.yml"}
}

func (d *DockerCompose) BuildCmd() (cmd *exec.Cmd) {
	args := []string{"-f", d.Path, "build"}
	if d.WorkDir != "" {
		args = append(args, "--project-directory", d.WorkDir)
	}
	cmd = exec.Command("docker-compose", args...)
	return
}

func (d *DockerCompose) Load() (err error) {
	// detect loaded
	if d.Compose.Version != "" {
		return
	}
	out, err := ioutil.ReadFile(d.Path)
	if err != nil {
		panic(fmt.Sprintf("Load docker-compose error, file: %s\nerror: %s", d.Path, err))
	}
	err = yaml.Unmarshal(out, &d.Compose)
	if err != nil {
		panic(fmt.Sprintf("Load docker-compose error, file: %s\nerror: %s", d.Path, err))
	}
	return
}

func (d *DockerCompose) Images() (images []common.Image, err error) {
	if err = d.Load(); err != nil {
		return
	}
	for name, service := range d.Compose.Services {
		if service.Build == "" {
			continue
		}
		if service.Image == "" {
			err = fmt.Errorf("put 'image: \"example\"' phase in %s to specific image name for service: '%s'", d.Path, name)
			return
		}
		images = append(images, fmt.Sprintf("%s:latest", service.Image))
	}
	return
}
