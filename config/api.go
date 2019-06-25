package config

import "time"

type api struct {
	Port        uint
	CORS        bool
	DBFile      string
	PluginsPath string
	Auth        auth
}

type auth struct {
	Secure            bool
	AuthExpiration    time.Duration
	RefreshExpiration time.Duration
}
