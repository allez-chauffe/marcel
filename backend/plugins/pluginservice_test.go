package plugins

import "testing"

func TestCheckExtension(t *testing.T) {
	if _, err := CheckExtension("marcel.zip"); err != nil {
		t.Error(err)
	}

	if _, err := CheckExtension("marcel.docx"); err == nil {
		t.Error(err)
	}
}
