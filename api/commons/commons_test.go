package commons

import (
	"testing"
)

func TestGetUID(t *testing.T) {
	const nbTests = 10
	var uids = [nbTests]string{}

	//generate a new UID and check whether it's unique
	for i := 0; i < nbTests; i++ {
		uid := GetUID()

		for j := 0; j < i; j++ {
			if uids[j] == uid {
				t.Error("UIDs should all be differents")
			}

		}

		uids[i] = uid
	}
}

func TestIsInArray(t *testing.T) {
	acceptedExtensions := []string{"zip", "gzip", "tar"}

	if r, _ := IsInArray("rar", acceptedExtensions); r == true {
		t.Fail()
	}

	if r, index := IsInArray("tar", acceptedExtensions); r == false || index != 2 {
		t.Fail()
	}
}

func TestBasename(t *testing.T) {
	if r := FileBasename("marcel.zip"); r != "marcel" {
		t.Errorf("Should get 'marcel', got %v", r)
	}

	if r := FileBasename("beliveau.tmp.zip"); r != "beliveau.tmp" {
		t.Errorf("Should get 'beliveau.tmp', got %v", r)
	}
}
