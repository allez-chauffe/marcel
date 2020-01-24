package plugins

import (
	"path/filepath"

	"github.com/Zenika/marcel/config"
)

type Plugin struct {
	Path         string       `json:"path"`
	Repositories []Repository `json:"versions"`
	Versions     []Version    `json:"versions"`
}

type Repository string

type Version struct {
	ID         string     `json:"id"`
	Repository Repository `json:"repository"`
	Version    string     `json:"version"`
	Short      string     `json:"short"`
	Long       string     `json:"long"`
	Widgets    []Widget   `json:"widgets"`
}

type Widget struct {
	Name  string `json:"name"`
	Short string `json:"short"`
	Long  string `json:"long"`
	Props []Prop `json:"props"`
}

type Prop struct {
	Name         string `json:"name"`
	Short        string `json:"short"`
	Long         string `json:"long"`
	Type         string `json:"type"`
	DefaultValue string `json:"value"`
}

// GetDirectory returns the plugin's static files directory path
func (p *Plugin) GetDirectory() string {
	// FIXME not sure this is a good idea to use config here
	return filepath.Join(config.Default().API().PluginsDir(), p.Path)
}
