package config

import "github.com/spf13/viper"

type Backoffice struct {
	v *viper.Viper
}

func (b Backoffice) Port() uint {
	return b.v.GetUint("backoffice.port")
}

func (b Backoffice) SetPort(p uint) {
	b.v.Set("backoffice.port", p)
}

func (b Backoffice) BasePath() string {
	return b.v.GetString("backoffice.basePath")
}

func (b Backoffice) SetBasePath(bp string) {
	b.v.Set("backoffice.basePath", bp)
}

func (b Backoffice) APIURI() string {
	return b.v.GetString("backoffice.apiURI")
}

func (b Backoffice) SetAPIURI(au string) {
	b.v.Set("backoffice.apiURI", au)
}

func (b Backoffice) FrontendURI() string {
	return b.v.GetString("backoffice.frontendURI")
}

func (b Backoffice) SetFrontendURI(fu string) {
	b.v.Set("backoffice.frontendURI", fu)
}
