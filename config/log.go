package config

import (
	"reflect"

	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

func registerLogLevelDecoder(dc *mapstructure.DecoderConfig) {
	dc.DecodeHook = mapstructure.ComposeDecodeHookFunc(
		dc.DecodeHook,
		func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
			if f.Kind() != reflect.String {
				return data, nil
			}
			if t != reflect.TypeOf(log.InfoLevel) {
				return data, nil
			}
			return log.ParseLevel(data.(string))
		},
	)
}
