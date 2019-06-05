package plugins

import (
	"path/filepath"

	"github.com/Zenika/MARCEL/config"
)

// Plugin represents a plugin configuration
//
// swagger:model
type Plugin struct {
	URL         string   `json:"url"`
	Versions    []string `json:"versions"`
	EltName     string   `json:"eltName"`
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

// GetDirectory returns the plugin's static files directory path
func (p *Plugin) GetDirectory() string {
	return filepath.Join(config.Config.PluginsPath, p.EltName)
}
