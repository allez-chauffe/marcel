package config

import "github.com/spf13/viper"

type Standalone viper.Viper

func (s *Standalone) viper() *viper.Viper {
	return (*viper.Viper)(s)
}

func (s *Standalone) Port() uint {
	return s.viper().GetUint("standalone.port")
}

func (s *Standalone) SetPort(p uint) {
	s.viper().Set("standalone.port", p)
}
