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
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Frontend    Frontend  `json:"frontend"`
	Backend     Backend   `json:"backend"`
}

func NewPlugin() (*Plugin) {
	var p = new(Plugin)

	p.Frontend = *NewFrontend()
	p.Backend = *NewBackend()

	return p
}

type Frontend struct {
	EltName   string `json:"eltName"`
	Cols      int `json:"cols"`
	Rows      int `json:"rows"`
	Fixedsize int `json:"fixedsize"`
	Props     []Props `json:"props"`
}

func NewFrontend() *Frontend {
	var f = new(Frontend)

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
