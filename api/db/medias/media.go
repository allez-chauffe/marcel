package medias

// Media represents a media configuration
type Media struct {
	ID          int                    `json:"id" boltholdKey:"ID"`
	Name        string                 `json:"name"`
	IsActive    bool                   `json:"isactive"`
	Description string                 `json:"description"`
	Rows        int                    `json:"rows"`
	Cols        int                    `json:"cols"`
	Owner       string                 `json:"owner"`
	ScreenRatio float64                `json:"screenRatio"`
	DisplayGrid bool                   `json:"displayGrid"`
	Stylesvar   map[string]interface{} `json:"stylesvar"`
	Plugins     []MediaPlugin          `json:"plugins"`
}

func New(owner string) *Media {
	return &Media{
		ID:          -1, // Let database autoincrement ID
		Owner:       owner,
		Rows:        10,
		Cols:        10,
		ScreenRatio: 16.0 / 9.0,
		DisplayGrid: true,
		Plugins:     []MediaPlugin{},
		Stylesvar:   make(map[string]interface{}),
	}
}

func (m *Media) GetID() interface{} {
	return m.ID
}

func (m *Media) SetID(id interface{}) {
	m.ID = id.(int)
}

// MediaPlugin represents a plugin configuration for the media
type MediaPlugin struct {
	InstanceID string               `json:"instanceId"`
	EltName    string               `json:"eltName"`
	FrontEnd   *MediaPluginFrontEnd `json:"frontend"`
}

type MediaPluginFrontEnd struct {
	X     int                    `json:"x"`
	Y     int                    `json:"y"`
	Rows  int                    `json:"rows"`
	Cols  int                    `json:"cols"`
	Props map[string]interface{} `json:"props"`
}
