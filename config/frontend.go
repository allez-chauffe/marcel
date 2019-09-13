package config

type privateFrontend struct {
	Port     uint
	BasePath string
	APIURI   string
}

type frontend struct {
	*privateFrontend
}
