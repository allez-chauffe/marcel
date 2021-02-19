package config

import (
	"github.com/spf13/viper"
)

type Frontend viper.Viper

func (f *Frontend) viper() *viper.Viper {
	return (*viper.Viper)(f)
}

func (f *Frontend) BasePath() string {
	return f.viper().GetString("frontend.basePath")
}

func (f *Frontend) SetBasePath(bp string) {
	f.viper().Set("frontend.basePath", bp)
}

func (f *Frontend) SetDefaults() {
	f.viper().SetDefault("frontend.basePath", "/front")
}
