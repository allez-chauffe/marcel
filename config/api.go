package config

import (
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

type API viper.Viper

func (a *API) viper() *viper.Viper {
	return (*viper.Viper)(a)
}

func (a *API) BasePath() string {
	return a.viper().GetString("api.basePath")
}

func (a *API) SetBasePath(bp string) {
	a.viper().Set("api.basePath", bp)
}

func (a *API) DBFile() string {
	return a.resolveDataDirPath(a.viper().GetString("api.dbFile"))
}

func (a *API) SetDBFile(df string) {
	a.viper().Set("api.dbFile", df)
}

func (a *API) PluginsDir() string {
	return a.resolveDataDirPath(a.viper().GetString("api.pluginsDir"))
}

func (a *API) SetPluginsDir(pd string) {
	a.viper().Set("api.pluginsDir", pd)
}

func (a *API) MediasDir() string {
	return a.resolveDataDirPath(a.viper().GetString("api.mediasDir"))
}

func (a *API) SetMediasDir(md string) {
	a.viper().Set("api.mediasDir", md)
}

func (a *API) DataDir() string {
	return a.viper().GetString("api.dataDir")
}

func (a *API) SetDataDir(dd string) {
	a.viper().Set("api.dataDir", dd)
}

func (a *API) resolveDataDirPath(pPath string) string {
	var path = os.ExpandEnv(pPath)
	if !filepath.IsAbs(path) {
		path = filepath.Join(os.ExpandEnv(a.DataDir()), path)
	}
	return filepath.Clean(path)
}

func (a *API) SetDefaults() {
	a.viper().SetDefault("api.basePath", "/api")
	a.viper().SetDefault("api.dbFile", "marcel.db")
	a.viper().SetDefault("api.pluginsDir", "plugins")
	a.viper().SetDefault("api.mediasDir", "medias")
	a.viper().SetDefault("api.dataDir", "")
	a.Auth().SetDefaults()
}

func (a *API) Auth() *Auth {
	return (*Auth)(a)
}

type Auth viper.Viper

func (a *Auth) viper() *viper.Viper {
	return (*viper.Viper)(a)
}

func (a *Auth) Secure() bool {
	return a.viper().GetBool("api.auth.secure")
}

func (a *Auth) SetSecure(s bool) {
	a.viper().Set("api.auth.secure", s)
}

func (a *Auth) Expiration() time.Duration {
	return a.viper().GetDuration("api.auth.expiration")
}

func (a *Auth) SetExpiration(e time.Duration) {
	a.viper().Set("api.auth.expiration", e)
}

func (a *Auth) RefreshExpiration() time.Duration {
	return a.viper().GetDuration("api.auth.refreshExpiration")
}

func (a *Auth) SetRefreshExpiration(re time.Duration) {
	a.viper().Set("api.auth.refreshExpiration", re)
}

func (a *Auth) SetDefaults() {
	a.viper().SetDefault("api.auth.secure", true)
	a.viper().SetDefault("api.auth.expiration", 8*time.Hour)
	a.viper().SetDefault("api.auth.refreshExpiration", 15*24*time.Hour)
}
