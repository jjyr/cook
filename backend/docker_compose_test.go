package backend

import (
	"testing"
	"strings"
	"runtime"
	"path"
	"fmt"
)

func projectDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Dir(path.Dir(filename))
}

func TestDockerCompose_Images(t *testing.T) {
	dc := NewDockerCompose()
	dc.Path = fmt.Sprintf("%s/test/docker-compose.yml", projectDir())
	images, err := dc.Images()
	if err != nil {
		t.Error(err)
	}
	for _, image := range images {
		if !strings.HasSuffix(image, ":latest") {
			t.Errorf("image %s is not tagged as latest", image)
		}
	}
}
