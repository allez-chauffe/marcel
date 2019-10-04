package config

import (
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func (c *ConfigType) FlagBool(flags *pflag.FlagSet, name string, value bool, usage string, keys ...string) error {
	flags.Bool(name, value, usage)
	return c.BindPFlag(flags, name, keys...)
}

func (c *ConfigType) FlagDuration(flags *pflag.FlagSet, name string, value time.Duration, usage string, keys ...string) error {
	flags.Duration(name, value, usage)
	return c.BindPFlag(flags, name, keys...)
}

func (c *ConfigType) FlagString(flags *pflag.FlagSet, name string, value string, usage string, keys ...string) error {
	flags.String(name, value, usage)
	return c.BindPFlag(flags, name, keys...)
}

func (c *ConfigType) FlagUint(flags *pflag.FlagSet, name string, value uint, usage string, keys ...string) error {
	flags.Uint(name, value, usage)
	return c.BindPFlag(flags, name, keys...)
}

func (c *ConfigType) FlagUintP(flags *pflag.FlagSet, name string, shorthand string, value uint, usage string, keys ...string) error {
	flags.UintP(name, shorthand, value, usage)
	return c.BindPFlag(flags, name, keys...)
}

func (c *ConfigType) BindPFlag(flags *pflag.FlagSet, name string, keys ...string) error {
	var cfg = (*viper.Viper)(c)
	var flag = flags.Lookup(name)
	for _, key := range keys {
		if err := cfg.BindPFlag(key, flag); err != nil {
			return err
		}
	}
	return nil
}
