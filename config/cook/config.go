package cook

import "github.com/jjyr/cook/common"

type CookConfig struct {
	Build struct {
		Dockerfiles []common.Dockerfile `yaml:"dockerfiles"`
	} `yaml:"build"`
	Target struct {
		Servers []common.Server `yaml:"servers"`
	} `yaml:"target"`
	Deploy struct {
		DockerComposes []common.DockerCompose `yaml:"docker_composes"`
	} `yaml:"deploy"`
}
