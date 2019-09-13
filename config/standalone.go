package config

type privateStandalone struct {
	Port uint
}

type standalone struct {
	*privateStandalone
}
