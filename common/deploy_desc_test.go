package common

import "testing"

func TestSetDeployDescDefault(t *testing.T) {
	d := DeployDesc{}
	if d == defaultDeployDesc {
		t.Errorf("initialed struct should not equal with defaultDeployDesc")
	}
	SetDeployDescDefault(&d)
	if d != defaultDeployDesc {
		t.Errorf("SetDeployDescDefault not work")
	}
}
