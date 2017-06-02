package media

type Media struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Config  MediaConfig `json:"config"`
	Plugins []MediaPlugins `json:"plugins"`
}

type MediaConfig struct {
	Styles []string `json:"styles"`
}

type MediaPlugins struct {
	Name       string `json:"name"`
	EltName    string `json:"eltName"`
	Files      []string `json:"files"`
	PropValues MediaPluginProps `json:"propValues"`
}

type MediaPluginProps struct {
	Values []interface{}
}
