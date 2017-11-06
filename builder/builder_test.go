package builder

import (
	"testing"
	"strings"
	"github.com/jjyr/cook/common"
)

func Test_getCurrentRef(t *testing.T) {
	ref, err := getCurrentRef()
	if err != nil {
		t.Fatal(err)
	}
	if ref == "" {
		t.Error("ref is blank")
	}
	if strings.ContainsAny(ref, "\r\n ") {
		t.Error("ref contains invalid chars, ref: %s", ref)
	}
}

func Test_buildCommand(t *testing.T) {
	b := NewBuilder(common.BuildDesc{})
	name, args := b.buildCommand("test")
	if name != "docker" || args[0] != "build" {
		t.Fatalf("command error, name:%s args:%+v", name, args)
	}

	b = NewBuilder(common.BuildDesc{Command: "echo 'customized'"})
	name, args = b.buildCommand("test")
	if name != "echo" || args[0] != "'customized'" {
		t.Error("customize command not work")
	}
}
