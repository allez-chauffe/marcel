package plugins

import (
	"fmt"
	"net/url"
	"path/filepath"

	"github.com/Zenika/marcel/config"
)

type Plugin struct {
	ID       string    `json:"id"`
	URL      string    `json:"url"`
	EltName  string    `json:"eltName"`
	Versions []Version `json:"versions"`
}

type Version struct {
	Version     string   `json:"version"`
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

// Path returns plugin's URL without scheme
func (p *Plugin) Path() (string, error) {
	if p.URL == "" {
		return p.EltName, nil
	}

	u, err := url.Parse(p.URL)
	if err != nil {
		return "", fmt.Errorf("Invalid plugin URL %s: %w", p.URL, err)
	}

	return u.Host + u.Path, nil
}

// GetDirectory returns the plugin's static files directory path
func (p *Plugin) GetDirectory() string {
	return filepath.Join(config.Default().API().PluginsDir(), p.ID) // FIXME wrong
}
