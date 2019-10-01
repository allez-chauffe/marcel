package config

import "github.com/spf13/viper"

type Standalone viper.Viper

func (s *Standalone) Port() uint {
	return s.cfg().GetUint("standalone.port")
}

func (s *Standalone) SetPort(p uint) {
	s.cfg().Set("standalone.port", p)
}

func (s *Standalone) cfg() *viper.Viper {
	return (*viper.Viper)(s)
}
