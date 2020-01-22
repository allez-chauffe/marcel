package plugins

import (
	"path/filepath"

	"github.com/Zenika/marcel/config"
)

type Plugin struct {
	ID       string             `json:"id"`
	Path     string             `json:"path" boltholdIndex:"Path"`
	URL      string             `json:"url"`
	Versions map[string]Version `json:"versions"`
}

type Version struct {
	Short   string   `json:"short"`
	Long    string   `json:"long"`
	Widgets []Widget `json:"widgets"`
}

type Widget struct {
	Name  string `json:"name"`
	Short string `json:"short"`
	Long  string `json:"long"`
	Cols  int    `json:"cols"` // FIXME Are these defaults cols/rows? min? max?
	Rows  int    `json:"rows"`
	Props []Prop `json:"props"`
}

type Prop struct {
	Name  string `json:"name"`
	Short string `json:"short"`
	Long  string `json:"long"`
	Type  string `json:"type"`
	Value string `json:"value"` // FIXME is this the default value?
}

// GetDirectory returns the plugin's static files directory path
func (p *Plugin) GetDirectory() string {
	// FIXME not sure this is a good idea to use config here
	return filepath.Join(config.Default().API().PluginsDir(), p.Path)
}
