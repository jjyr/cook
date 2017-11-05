package builder

import "github.com/jjyr/cook/common"

type Builder struct {
}

func NewBuilder(dockerfile common.Dockerfile) (*Builder) {
	return &Builder{}
}

func (b *Builder) Build() (err error) {
	return
}
