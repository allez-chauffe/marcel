package medias

// Media represents a media configuration
type Media struct {
	ID          int                    `json:"id"`
	Short       string                 `json:"short"`
	Long        string                 `json:"long"`
	Owner       string                 `json:"owner"`
	IsActive    bool                   `json:"isactive"`
	Rows        int                    `json:"rows"`
	Cols        int                    `json:"cols"`
	ScreenRatio float64                `json:"screenRatio"` // FIXME rename to ratio?
	Stylesvar   map[string]interface{} `json:"stylesvar"`
	DisplayGrid bool                   `json:"displayGrid"`
	Widgets     []Widget               `json:"widgets"`
}

func New(owner string) *Media {
	return &Media{
		Owner:       owner,
		Rows:        10,
		Cols:        10,
		ScreenRatio: 16.0 / 9.0,
		DisplayGrid: true,
		Stylesvar:   make(map[string]interface{}),
		Widgets:     []Widget{},
	}
}

type Widget struct {
	PluginID   string                 `json:"pluginId"`
	WidgetName string                 `json:"widgetName"`
	Short      string                 `json:"short"`
	Long       string                 `json:"long"`
	X          int                    `json:"x"`
	Y          int                    `json:"y"`
	Rows       int                    `json:"rows"`
	Cols       int                    `json:"cols"`
	Props      map[string]interface{} `json:"props"`
}
