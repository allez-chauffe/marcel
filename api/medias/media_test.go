package medias

import "testing"

func TestNewConfiguration(t *testing.T) {
	c := NewConfiguration()

	if len(c.Medias) != 0 {
		t.Error("New configuration should have a medias list of 0 element")
	}

	if c.LastID != 0 {
		t.Error("A new medias configuration should have a 'LastId == 0'")
	}
}

func TestNewMedia(t *testing.T) {
	m := NewMedia()

	if len(m.Plugins) != 0 {
		t.Error("A new media should have an empty list of plugins")
	}

	if m.Description != "" {
		t.Error("A new media should have a description equals to \"\"")
	}

	if m.Name != "" {
		t.Error("A new media should have a name equals to \"\"")
	}

	if m.Rows != 0 {
		t.Error("A new media should have Rows==0")
	}
	if m.Cols != 0 {
		t.Error("A new media should have Cols==0")
	}
	if m.ID != 0 {
		t.Error("A new media should have ID==0")
	}

}
