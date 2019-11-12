package config

import "github.com/spf13/viper"

type Backoffice viper.Viper

func (b *Backoffice) viper() *viper.Viper {
	return (*viper.Viper)(b)
}

func (b *Backoffice) BasePath() string {
	return b.viper().GetString("backoffice.basePath")
}

func (b *Backoffice) SetBasePath(bp string) {
	b.viper().Set("backoffice.basePath", bp)
}

func (b *Backoffice) APIURI() string {
	return b.viper().GetString("backoffice.apiURI")
}

func (b *Backoffice) SetAPIURI(au string) {
	b.viper().Set("backoffice.apiURI", au)
}

func (b *Backoffice) FrontendURI() string {
	return b.viper().GetString("backoffice.frontendURI")
}

func (b *Backoffice) SetFrontendURI(fu string) {
	b.viper().Set("backoffice.frontendURI", fu)
}

func (b *Backoffice) SetDefaults() {
	b.viper().SetDefault("backoffice.basePath", "/")
	b.viper().SetDefault("backoffice.apiURI", "/api")
	b.viper().SetDefault("backoffice.frontendURI", "/front")
}
