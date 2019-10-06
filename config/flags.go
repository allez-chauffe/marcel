package config

import (
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func (c *Config) FlagBool(flags *pflag.FlagSet, name string, value bool, usage string, keys ...string) (*bool, error) {
	return flags.Bool(name, value, usage), c.BindPFlag(flags, name, keys...)
}

func (c *Config) FlagDuration(flags *pflag.FlagSet, name string, value time.Duration, usage string, keys ...string) (*time.Duration, error) {
	return flags.Duration(name, value, usage), c.BindPFlag(flags, name, keys...)
}

func (c *Config) FlagString(flags *pflag.FlagSet, name string, value string, usage string, keys ...string) (*string, error) {
	return flags.String(name, value, usage), c.BindPFlag(flags, name, keys...)
}

func (c *Config) FlagUint(flags *pflag.FlagSet, name string, value uint, usage string, keys ...string) (*uint, error) {
	return flags.Uint(name, value, usage), c.BindPFlag(flags, name, keys...)
}

func (c *Config) FlagUintP(flags *pflag.FlagSet, name string, shorthand string, value uint, usage string, keys ...string) (*uint, error) {
	return flags.UintP(name, shorthand, value, usage), c.BindPFlag(flags, name, keys...)
}

func (c *Config) BindPFlag(flags *pflag.FlagSet, name string, keys ...string) error {
	var cfg = (*viper.Viper)(c)
	var flag = flags.Lookup(name)
	for _, key := range keys {
		if err := cfg.BindPFlag(key, flag); err != nil {
			return err
		}
	}
	return nil
}
