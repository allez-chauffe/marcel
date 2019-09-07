package config

import "time"

type api struct {
	Port       uint
	BasePath   string
	CORS       bool
	DBFile     string
	PluginsDir string
	MediasDir  string
	Auth       auth
}

type auth struct {
	Secure            bool
	Expiration        time.Duration
	RefreshExpiration time.Duration
}
