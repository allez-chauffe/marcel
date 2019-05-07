package config

import "time"

type auth struct {
	Secured           bool
	AuthExpiration    time.Duration
	RefreshExpiration time.Duration
}
