package plugins

// PluginsConfiguration encapsulates all configuration data
//
// swagger:model
type Configuration struct {
	Plugins []Plugin `json:"plugins"`
}

func NewConfiguration() *Configuration {
	var configuration = new(Configuration)

	configuration.Plugins = []Plugin{}

	return configuration
}

// Plugin represents a plugin configuration
//
// swagger:model
type Plugin struct {
	EltName     string     `json:"eltName"`
	Description string     `json:"description"`
	Frontend    Frontend   `json:"frontend"`
	Backend     Backend    `json:"backend"`
}

func NewPlugin() (*Plugin) {
	var p = new(Plugin)

	p.Frontend = *NewFrontend()
	p.Backend = *NewBackend()

	return p
}

type Frontend struct {
	Cols  int `json:"cols"`
	Rows  int `json:"rows"`
	Props []Props `json:"props"`
}

func NewFrontend() *Frontend {
	var f = new(Frontend)

	f.Cols = 0
	f.Rows = 0
	f.Props = []Props{}

	return f
}

type Props struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Value       string `json:"value"`
}

type Backend struct {
	Ports []int                  `json:"ports"`
	Props map[string]interface{} `json:"props"`
}

func NewBackend() *Backend {
	var b = new(Backend)

	b.Ports = []int{}
	b.Props = make(map[string]interface{})

	return b
}
