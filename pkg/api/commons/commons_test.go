package commons

import (
	"testing"
)

func TestIsInArray(t *testing.T) {
	acceptedExtensions := []string{"zip", "gzip", "tar"}

	if r, _ := IsInArray("rar", acceptedExtensions); r == true {
		t.Fail()
	}

	if r, index := IsInArray("tar", acceptedExtensions); r == false || index != 2 {
		t.Fail()
	}
}
