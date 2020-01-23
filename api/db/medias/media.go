package medias

type Media struct {
	ID       string            `json:"id"`
	Short    string            `json:"short"`
	Long     string            `json:"long"`
	Owner    string            `json:"owner"`
	IsActive bool              `json:"isactive"`
	Rows     int               `json:"rows"`
	Cols     int               `json:"cols"`
	Ratio    float64           `json:"ratio"`
	Style    map[string]string `json:"style"`
	Widgets  []Widget          `json:"widgets"`
	Grid     []GridItem        `json:"grid"`
}

func New(owner string) *Media {
	return &Media{
		Owner:   owner,
		Rows:    10,
		Cols:    10,
		Ratio:   16.0 / 9.0,
		Style:   make(map[string]string),
		Widgets: []Widget{},
	}
}

type Widget struct {
	ID            string                 `json:"id"`
	PluginID      string                 `json:"pluginId"`
	PluginVersion string                 `json:"pluginVersion"`
	WidgetName    string                 `json:"widgetName"`
	Short         string                 `json:"short"`
	Long          string                 `json:"long"`
	Props         map[string]interface{} `json:"props"`
}

type GridItem struct {
	WidgetID string `widgetId`
	X        int    `json:"x"`
	Y        int    `json:"y"`
	Rows     int    `json:"rows"`
	Cols     int    `json:"cols"`
}
