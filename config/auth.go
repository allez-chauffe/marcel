package config

import "time"

type auth struct {
	Secure            bool
	AuthExpiration    time.Duration
	RefreshExpiration time.Duration
}
