package config

import "github.com/spf13/viper"

type Frontend viper.Viper

func (f *Frontend) Port() uint {
	return f.cfg().GetUint("frontend.port")
}

func (f *Frontend) SetPort(p uint) {
	f.cfg().Set("frontend.port", p)
}

func (f *Frontend) BasePath() string {
	return f.cfg().GetString("frontend.basePath")
}

func (f *Frontend) SetBasePath(bp string) {
	f.cfg().Set("frontend.basePath", bp)
}

func (f *Frontend) APIURI() string {
	return f.cfg().GetString("frontend.apiURI")
}

func (f *Frontend) SetAPIURI(au string) {
	f.cfg().Set("frontend.apiURI", au)
}

func (f *Frontend) cfg() *viper.Viper {
	return (*viper.Viper)(f)
}
