package medias

// Media represents a media configuration
//
// swagger:model
type Media struct {
	// the id for this media
	//
	// required: true
	// unique: true
	// min: 1
	ID          int                    `json:"id"`
	Name        string                 `json:"name"`
	IsActive    bool                   `json:"isactive"`
	Description string                 `json:"description"`
	Rows        int                    `json:"rows"`
	Cols        int                    `json:"cols"`
	Stylesvar   map[string]interface{} `json:"stylesvar"`
	Plugins     []MediaPlugin          `json:"plugins"`
	Owner       string                 `json:"owner"`
	ScreenRatio float64                `json:"screenRatio"`
	DisplayGrid bool                   `json:"displayGrid"`
}

func NewMedia() *Media {
	var media = new(Media)

	media.Stylesvar = make(map[string]interface{})
	media.Plugins = []MediaPlugin{}
	media.Rows = 10
	media.Cols = 10
	media.ScreenRatio = 16.0 / 9.0
	media.DisplayGrid = true

	return media
}

// MediaPlugin represents a plugin configuration for the media
//
// Properties and configuration for a plugin used in the media
//
// swagger:model
type MediaPlugin struct {
	InstanceId string               `json:"instanceId"`
	EltName    string               `json:"eltName"`
	FrontEnd   *MediaPluginFrontEnd `json:"frontend"`
	BackEnd    *MediaPluginBackEnd  `json:"backend"`
}

type MediaPluginFrontEnd struct {
	//Files []string               `json:"files"`
	X     int                    `json:"x"`
	Y     int                    `json:"y"`
	Rows  int                    `json:"rows"`
	Cols  int                    `json:"cols"`
	Props map[string]interface{} `json:"props"`
}

type MediaPluginBackEnd struct {
	Port              int                    `json:"port"`
	Props             map[string]interface{} `json:"props"`
	DockerImageName   string
	DockerContainerId string
}
