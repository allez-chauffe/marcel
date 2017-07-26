package commons

import (
	"testing"
)

func TestGetUID(t *testing.T) {
	t.Log("GetUID test")
	const nbTests = 10
	var uids = [nbTests]string{}

	//generate a new UID and check whether it's unique
	for i := 0; i < nbTests; i++ {
		uid := GetUID()

		for j:=0; j<i; j++ {
			if uids[j] == uid {
				t.Error("UIDs should all be differents")
			}

		}

		uids[i] = uid
	}
}
