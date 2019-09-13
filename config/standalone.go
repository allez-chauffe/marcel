package config

import "github.com/spf13/viper"

type Standalone struct {
	v *viper.Viper
}

func (s Standalone) Port() uint {
	return s.v.GetUint("standalone.port")
}

func (s Standalone) SetPort(p uint) {
	s.v.Set("standalone.port", p)
}
