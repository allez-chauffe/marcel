package config

import "time"

type auth struct {
	Port              uint
	Secured           bool
	AuthExpiration    time.Duration
	RefreshExpiration time.Duration
	Domain            string
	BaseURL           string
	UsersFile         string //FIXME to be removed
}
