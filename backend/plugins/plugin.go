package plugins

// PluginsConfiguration encapsulates all configuration data
//
// swagger:model
type Configuration struct {
	LastID  int     `json:"lastid"`
	Plugins []Plugin `json:"plugins"`
}

func NewConfiguration() *Configuration {
	var configuration = new(Configuration)

	configuration.LastID = 0
	configuration.Plugins = []Plugin{}

	return configuration
}

// Plugin represents a plugin configuration
//
// swagger:model
type Plugin struct {
	ID          int    `json:"id"`
	Name        string        `json:"title"`
	Description string    `json:"description"`
	Frontend    []Frontend `json:"frontend"`
	Backend     []Backend `json:"backend"`
}

func NewPlugin() (*Plugin) {
	var p = new(Plugin)

	return p
}

type Frontend struct {
	Key   string
	Value string
}

func NewFrontend() *Frontend {
	var p = new(Frontend)

	return p
}

type Backend struct {
	Key   string
	Value string
}

func NewBackend() *Backend {
	var p = new(Backend)

	return p
}
