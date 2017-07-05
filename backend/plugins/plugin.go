package plugins

// Plugin represents a plugin configuration
//
// swagger:model
type Plugin struct {
	ID     int           `json:"id"`
	Name   string        `json:"title"`
	Params []PluginParam `json:"params"`
}

type PluginParam struct {
	Key   string
	Value string
}
