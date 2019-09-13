package config

import "github.com/spf13/viper"

type Frontend struct {
	v *viper.Viper
}

func (f Frontend) Port() uint {
	return f.v.GetUint("frontend.port")
}

func (f Frontend) SetPort(p uint) {
	f.v.Set("frontend.port", p)
}

func (f Frontend) BasePath() string {
	return f.v.GetString("frontend.basePath")
}

func (f Frontend) SetBasePath(bp string) {
	f.v.Set("frontend.basePath", bp)
}

func (f Frontend) APIURI() string {
	return f.v.GetString("frontend.apiURI")
}

func (f Frontend) SetAPIURI(au string) {
	f.v.Set("frontend.apiURI", au)
}
