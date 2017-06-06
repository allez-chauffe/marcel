package media

/**
The global attributes for a Media
 */
type Media struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Config  MediaConfig `json:"config"`
	Plugins []MediaPlugin `json:"plugins"`
}

type MediaConfig struct {
	Styles []string `json:"styles"`
}

/**
Properties and configuration for a plugin used in the media
 */
type MediaPlugin struct {
	Name       string `json:"name"`
	EltName    string `json:"eltName"`
	Files      []string `json:"files"`
	PropValues MediaPluginProps `json:"propValues"`
}

/**
Because we don't know what will compounds the props for a plugin, we use a map[string] interface{}
 */
type MediaPluginProps struct {
	X map[string]interface{} `json:"-"`
}
