package common

type DeployDesc struct {
	Type        string `yaml:"type"`
	Path        string `yaml:"path"`
	WorkDir     string `yaml:"work_dir"`
	ProjectName string `yaml:"project_name"`
}

var defaultDeployDesc = DeployDesc{Type: "docker-compose"}

func SetDeployDescDefault(desc *DeployDesc) {
	if desc.Type == "" {
		desc.Type = defaultDeployDesc.Type
	}
}
