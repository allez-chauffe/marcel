package plugins

import "testing"

func TestNewBackend(t *testing.T) {
	t.Log("NewBackend test")
	b := NewBackend()

	if len(b.Ports) != 0 {
		t.Error("New backend should have an array of ports of length 0")
	}
}

func TestNewFrontend(t *testing.T) {
	t.Log("NewFrontend test")
	f := NewFrontend()

	if f.Cols != 0 || f.Rows != 0 {
		t.Error("New frontend should get Cols and Rows set to 0")
	}
}

func TestNewPlugin(t *testing.T) {
	t.Log("NewPlugin test")
	p := NewPlugin()

	if p.EltName != "" {
		t.Error("EltName should be equal to \"\"")
	}

	if p.Name != "" {
		t.Error("Name should be equal to \"\"")
	}

	if p.Description != "" {
		t.Error("Description should be equal to \"\"")
	}
}

func TestNewConfiguration(t *testing.T) {
	t.Log("NewConfiguration test")
	c := NewConfiguration()

	if len(c.Plugins) != 0 {
		t.Error("New configuration should have a plugins list of 0 element")
	}
}