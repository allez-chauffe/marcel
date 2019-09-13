package config

type privateBackoffice struct {
	Port        uint
	BasePath    string
	APIURI      string
	FrontendURI string
}

type backoffice struct {
	*privateBackoffice
}
