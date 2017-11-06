package builder

import (
	"github.com/jjyr/cook/common"
	"os/exec"
	"strings"
	"fmt"
)

type Builder struct {
	BuildDesc common.BuildDesc
}

func NewBuilder(buildDesc common.BuildDesc) (*Builder) {
	return &Builder{BuildDesc: buildDesc}
}

func getCurrentBranch() (branch string, err error) {
	cmd := exec.Command("git", "branch", "-q")
	output, err := cmd.Output()
	if err != nil {
		return
	}
	outputBranch := strings.SplitN(string(output), "\n", 1)[0]
	outputBranch = strings.Trim(outputBranch, "\n\r ")
	if !strings.HasPrefix(outputBranch, "* ") {
		err = fmt.Errorf("uncognized git branch '%s', command:\ngit branch -q\noutput:\n%s", outputBranch, string(output))
		return
	}
	branch = outputBranch[2:]
	return
}

func getCurrentCommitID() (commitID string, err error) {
	cmd := exec.Command("git", "rev-parse", "--short", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return
	}
	commitID = strings.Trim(string(output), "\n\r ")
	if commitID == "" {
		err = fmt.Errorf("uncognized git commitID, command:\ngit rev-parse --short HEAD\noutput:\n%s", string(output))
	}
	return
}

func getCurrentRef() (ref string, err error) {
	branch, err := getCurrentBranch()
	if err != nil {
		return
	}
	commitID, err := getCurrentCommitID()
	if err != nil {
		return
	}
	ref = fmt.Sprintf("%s-%s", branch, commitID)
	return
}

func (b *Builder) buildCommand(tag string) (name string, args []string) {
	if b.BuildDesc.Command == "" {
		path := b.BuildDesc.Path
		if path == "" {
			path = "."
		}
		// default command
		tagName := fmt.Sprintf("%s:%s", b.BuildDesc.Image, tag)
		name = "docker"
		args = []string{"build", "-t", tagName, path}
		return
	}

	commands := strings.Split(b.BuildDesc.Command, " ")
	name = commands[0]
	args = commands[1:]
	return
}

func (b *Builder) Build() (err error) {
	ref, err := getCurrentRef()
	if err != nil {
		return
	}
	command, args := b.buildCommand(ref)
	cmd := exec.Command(command, args...)
	cmd.Dir = b.BuildDesc.BuildDir
	err = cmd.Run()
	return
}
