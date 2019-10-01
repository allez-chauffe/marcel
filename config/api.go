package config

import (
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

type API viper.Viper

func (a *API) cfg() *viper.Viper {
	return (*viper.Viper)(a)
}

func (a *API) Port() uint {
	return a.cfg().GetUint("api.port")
}

func (a *API) SetPort(p uint) {
	a.cfg().Set("api.port", p)
}

func (a *API) BasePath() string {
	return a.cfg().GetString("api.basePath")
}

func (a *API) SetBasePath(bp string) {
	a.cfg().Set("api.basePath", bp)
}

func (a *API) CORS() bool {
	return a.cfg().GetBool("api.cors")
}

func (a *API) SetCORS(c bool) {
	a.cfg().Set("api.cors", c)
}

func (a *API) DBFile() string {
	return a.resolveDataDirPath(a.cfg().GetString("api.dbFile"))
}

func (a *API) SetDBFile(df string) {
	a.cfg().Set("api.dbFile", df)
}

func (a *API) PluginsDir() string {
	return a.resolveDataDirPath(a.cfg().GetString("api.pluginsDir"))
}

func (a *API) SetPluginsDir(pd string) {
	a.cfg().Set("api.pluginsDir", pd)
}

func (a *API) MediasDir() string {
	return a.resolveDataDirPath(a.cfg().GetString("api.mediasDir"))
}

func (a *API) SetMediasDir(pd string) {
	a.cfg().Set("api.mediasDir", pd)
}

func (a *API) DataDir() string {
	return a.cfg().GetString("api.dataDir")
}

func (a *API) SetDataDir(pd string) {
	a.cfg().Set("api.dataDir", pd)
}

func (a *API) resolveDataDirPath(pPath string) string {
	var path = os.ExpandEnv(pPath)
	if !filepath.IsAbs(path) {
		path = filepath.Join(os.ExpandEnv(a.DataDir()), path)
	}
	return filepath.Clean(path)
}

func (a *API) Auth() *Auth {
	return (*Auth)(a)
}

type Auth viper.Viper

func (a *Auth) Secure() bool {
	return a.cfg().GetBool("api.auth.secure")
}

func (a *Auth) SetSecure(s bool) {
	a.cfg().Set("api.auth.secure", s)
}

func (a *Auth) Expiration() time.Duration {
	return a.cfg().GetDuration("api.auth.expiration")
}

func (a *Auth) SetExpiration(e bool) {
	a.cfg().Set("api.auth.expiration", e)
}

func (a *Auth) RefreshExpiration() time.Duration {
	return a.cfg().GetDuration("api.auth.refreshExpiration")
}

func (a *Auth) SetRefreshExpiration(re bool) {
	a.cfg().Set("api.auth.refreshExpiration", re)
}

func (a *Auth) cfg() *viper.Viper {
	return (*viper.Viper)(a)
}
