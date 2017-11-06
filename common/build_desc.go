package common

type BuildDesc struct {
	Path     string `yaml:"path"`
	BuildDir string `yaml:"build_dir"`
	Image    Image  `yaml:"image"`
	Command  string `yaml:"command"`
}
