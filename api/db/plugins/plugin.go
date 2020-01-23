package plugins

import (
	"path/filepath"

	"github.com/Zenika/marcel/config"
)

type Plugin struct {
	ID       string    `json:"id"`
	Path     string    `json:"path" boltholdIndex:"Path"`
	URL      string    `json:"url"`
	Versions []Version `json:"versions"`
}

type Version struct {
	Version string   `json:"version"`
	Short   string   `json:"short"`
	Long    string   `json:"long"`
	Widgets []Widget `json:"widgets"`
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
