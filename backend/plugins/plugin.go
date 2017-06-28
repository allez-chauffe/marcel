package plugins

type Plugin struct {
	ID     int           `json:"id"`
	Name   string        `json:"title"`
	Params []PluginParam `json:"params"`
}

type PluginParam struct {
	Key   string
	Value string
}
