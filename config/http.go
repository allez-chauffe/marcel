package config

import "github.com/spf13/viper"

type HTTP viper.Viper

func (h *HTTP) viper() *viper.Viper {
	return (*viper.Viper)(h)
}

func (h *HTTP) Port() uint {
	return h.viper().GetUint("http.port")
}

func (h *HTTP) SetPort(p uint) {
	h.viper().Set("http.port", p)
}

func (h *HTTP) CORS() bool {
	return h.viper().GetBool("http.cors")
}

func (h *HTTP) SetCORS(c bool) {
	h.viper().Set("http.cors", c)
}

func (h *HTTP) SetDefaults() {
	h.viper().SetDefault("http.port", 8090)
	h.viper().SetDefault("http.cors", false)
}
