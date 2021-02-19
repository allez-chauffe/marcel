package xos

import (
	"os"

	"github.com/mattn/go-isatty"
)

// IsInteractive tests whether Stdin and Stdout are interactive.
func IsInteractive() bool {
	return isatty.IsTerminal(os.Stdin.Fd()) && isatty.IsTerminal(os.Stdout.Fd())
}
