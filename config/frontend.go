package config

import (
	"path"

	"github.com/spf13/viper"
)

type Frontend viper.Viper

func (f *Frontend) viper() *viper.Viper {
	return (*viper.Viper)(f)
}

func (f *Frontend) BasePath() string {
	return f.viper().GetString("frontend.basePath")
}

func (f *Frontend) AbsoluteBasePath() string {
	return path.Join((*HTTP)(f.viper()).BasePath(), f.BasePath())
}

func (f *Frontend) SetBasePath(bp string) {
	f.viper().Set("frontend.basePath", bp)
}

func (f *Frontend) SetDefaults() {
	f.viper().SetDefault("frontend.basePath", "/front")
}
