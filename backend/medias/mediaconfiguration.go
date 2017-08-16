package medias

// MediasConfig encapsulates all configuration data
//
// swagger:model
type Configuration struct {
	LastID int     `json:"lastid"`
	Medias []Media `json:"medias"`

	//PortsPool is an array of ports that were used by backends, but now are free (because of a deactivation)
	PortsPool []int
	//LastPortUsed is a counter and allow to generate a new free port number
	NextFreePortNumber int
}

func NewConfiguration() *Configuration {
	var configuration = new(Configuration)

	configuration.LastID = 0
	configuration.Medias = []Media{}

	configuration.NextFreePortNumber = 8100

	return configuration
}
