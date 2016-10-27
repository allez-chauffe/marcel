package weather

import (
	"testing"
)

func TestGetLocation(t *testing.T) {
	t.Parallel()

	location, err := getLocation()
	if err != nil {
		t.Error(err)
	}

	if location.City == nil {
		t.Error("Location's city is nil")
	}
	t.Log(location.City)
	t.Log(location.Status)
}
