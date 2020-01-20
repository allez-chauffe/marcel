package config

import (
	"github.com/spf13/viper"
)

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

func (b *Backoffice) SetDefaults() {
	b.viper().SetDefault("backoffice.basePath", "/")
}
