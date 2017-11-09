package backend

import "github.com/jjyr/cook/common"

type DeployBackend = *DockerCompose

func GetBackend(d common.DeployDesc) (db DeployBackend) {
	if d.Type != "docker-compose" {
		panic("unknown deploy type: " + d.Type)
	}
	db = NewDockerCompose()
	db.Path = d.Path
	db.WorkDir = d.WorkDir
	return
}
