package common

import "testing"

func TestSetServerDefault(t *testing.T) {
	s := Server{}
	if s == defaultServer {
		t.Errorf("initialed struct should not equal with defaultServer")
	}
	SetServerDefault(&s)
	if s != defaultServer {
		t.Errorf("SetServerDefault not work")
	}
}
