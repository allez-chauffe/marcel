package osutil

import "os/exec"

func Open(path string) error {
	return exec.Command("start", path).Start()
}
