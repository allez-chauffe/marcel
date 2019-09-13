package config

import "github.com/spf13/viper"

type Backoffice viper.Viper

func (b *Backoffice) Port() uint {
	return b.cfg().GetUint("backoffice.port")
}

func (b *Backoffice) SetPort(p uint) {
	b.cfg().Set("backoffice.port", p)
}

func (b *Backoffice) BasePath() string {
	return b.cfg().GetString("backoffice.basePath")
}

func (b *Backoffice) SetBasePath(bp string) {
	b.cfg().Set("backoffice.basePath", bp)
}

func (b *Backoffice) APIURI() string {
	return b.cfg().GetString("backoffice.apiURI")
}

func (b *Backoffice) SetAPIURI(au string) {
	b.cfg().Set("backoffice.apiURI", au)
}

func (b *Backoffice) FrontendURI() string {
	return b.cfg().GetString("backoffice.frontendURI")
}

func (b *Backoffice) SetFrontendURI(fu string) {
	b.cfg().Set("backoffice.frontendURI", fu)
}

func (b *Backoffice) cfg() *viper.Viper {
	return (*viper.Viper)(b)
}
