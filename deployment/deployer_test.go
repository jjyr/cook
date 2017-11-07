package deployment

import (
	"testing"
	"github.com/jjyr/cook/common"
)

func TestDeployer_Prepare(t *testing.T) {
	deploy := NewDeployer(common.Server{})
	if err := deploy.Prepare(""); err != nil {
		t.Error(err)
	}
}
