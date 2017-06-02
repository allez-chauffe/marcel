package media

type Media struct {
	ID     int    `json:"id"`
	Name   string `json:"title"`
	Params []MediaParam `json:"params"`
}

type MediaParam struct {
	Key   string
	Value string
}