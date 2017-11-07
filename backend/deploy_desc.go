package backend

type DeployDesc struct {
	Type    string `yaml:"type"`
	Path    string `yaml:"path"`
	WorkDir string `yaml:"work_dir"`
}

func (d *DeployDesc) GetBackend() (db DeployBackend) {
	if d.Type != "docker_compose" {
		panic("unknown deploy type: " + d.Type)
	}
	db = NewDockerCompose()
	db.Path = d.Path
	db.WorkDir = d.WorkDir
	return
}
