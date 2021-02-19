package plugins

import (
	"path/filepath"

	"github.com/allez-chauffe/marcel/pkg/config"
)

// Plugin represents a plugin configuration
type Plugin struct {
	URL      string   `json:"url"`
	Versions []string `json:"versions"`
	// FIXME: Do better to handle differnent primary key names
	EltName     string   `json:"eltName" boltholdKey:"EltName"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Frontend    Frontend `json:"frontend"`
}

type Frontend struct {
	Cols  int    `json:"cols"`
	Rows  int    `json:"rows"`
	Props []Prop `json:"props"`
}

type Prop struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Value       string `json:"value"`
}

func New(eltName string) *Plugin {
	return &Plugin{EltName: eltName}
}

// GetDirectory returns the plugin's static files directory path
func (p *Plugin) GetDirectory() string {
	return filepath.Join(config.Default().API().PluginsDir(), p.EltName)
}

func (p *Plugin) GetID() interface{} {
	return p.EltName
}

func (p *Plugin) SetID(id interface{}) {
	p.EltName = id.(string)
}
