package osutil

import (
	"os/exec"
)

func Open(path string) error {
	return exec.Command("xdg-open", path).Start()
}
