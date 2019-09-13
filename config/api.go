package config

import (
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

type API struct {
	v *viper.Viper
}

func (a API) Port() uint {
	return a.v.GetUint("api.port")
}

func (a API) SetPort(p uint) {
	a.v.Set("api.port", p)
}

func (a API) BasePath() string {
	return a.v.GetString("api.basePath")
}

func (a API) SetBasePath(bp string) {
	a.v.Set("api.basePath", bp)
}

func (a API) CORS() bool {
	return a.v.GetBool("api.cors")
}

func (a API) SetCORS(c bool) {
	a.v.Set("api.cors", c)
}

func (a API) DBFile() string {
	return a.resolveDataDirPath(a.v.GetString("api.dbFile"))
}

func (a API) SetDBFile(df string) {
	a.v.Set("api.dbFile", df)
}

func (a API) PluginsDir() string {
	return a.resolveDataDirPath(a.v.GetString("api.pluginsDir"))
}

func (a API) SetPluginsDir(pd string) {
	a.v.Set("api.pluginsDir", pd)
}

func (a API) MediasDir() string {
	return a.resolveDataDirPath(a.v.GetString("api.mediasDir"))
}

func (a API) SetMediasDir(pd string) {
	a.v.Set("api.mediasDir", pd)
}

func (a API) DataDir() string {
	return a.v.GetString("api.dataDir")
}

func (a API) SetDataDir(pd string) {
	a.v.Set("api.dataDir", pd)
}

func (a API) resolveDataDirPath(pPath string) string {
	var path = os.ExpandEnv(pPath)
	if !filepath.IsAbs(path) {
		path = filepath.Join(os.ExpandEnv(a.DataDir()), path)
	}
	return filepath.Clean(path)
}

func (a API) Auth() Auth {
	return Auth{a.v}
}

type Auth struct {
	v *viper.Viper
}

func (a Auth) Secure() bool {
	return a.v.GetBool("api.auth.secure")
}

func (a Auth) SetSecure(s bool) {
	a.v.Set("api.auth.secure", s)
}

func (a Auth) Expiration() time.Duration {
	return a.v.GetDuration("api.auth.expiration")
}

func (a Auth) SetExpiration(e bool) {
	a.v.Set("api.auth.expiration", e)
}

func (a Auth) RefreshExpiration() time.Duration {
	return a.v.GetDuration("api.auth.refreshExpiration")
}

func (a Auth) SetRefreshExpiration(re bool) {
	a.v.Set("api.auth.refreshExpiration", re)
}
