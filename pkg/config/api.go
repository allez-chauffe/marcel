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
	a.viper().SetDefault("api.pluginsDir", "plugins")
	a.viper().SetDefault("api.mediasDir", "medias")
	a.viper().SetDefault("api.dataDir", "")
	a.Auth().SetDefaults()
	a.DB().SetDefault()
}

func (a *API) Auth() *Auth {
	return (*Auth)(a)
}

func (a *API) DB() *DB {
	return (*DB)(a)
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

type DB viper.Viper

func (db *DB) SetDefault() {
	db.viper().SetDefault("api.db.driver", "bolt")
	db.Bolt().SetDefault()
	db.Postgres().SetDefault()
}

func (db *DB) viper() *viper.Viper {
	return (*viper.Viper)(db)
}

func (db *DB) Driver() string {
	return db.viper().GetString("api.db.driver")
}

func (db *DB) SetDriver(driver string) {
	db.viper().Set("api.db.driver", driver)
}

type Bolt viper.Viper

func (db *DB) Bolt() *Bolt {
	return (*Bolt)(db)
}

func (b *Bolt) SetDefault() {
	b.viper().SetDefault("api.db.bolt.file", "marcel.db")
}

func (b *Bolt) viper() *viper.Viper {
	return (*viper.Viper)(b)
}

func (b *Bolt) File() string {
	return (*API)(b).resolveDataDirPath(b.viper().GetString("api.db.bolt.file"))
}

func (b *Bolt) SetFile(file string) {
	b.viper().Set("api.db.bolt.file", file)
}

type Postgres viper.Viper

func (db *DB) Postgres() *Postgres {
	return (*Postgres)(db)
}

func (p *Postgres) SetDefault() {
	p.viper().SetDefault("api.db.postgres.dbName", "postgres")
	p.viper().SetDefault("api.db.postgres.host", "localhost")
	p.viper().SetDefault("api.db.postgres.port", 5432)
}

func (p *Postgres) viper() *viper.Viper {
	return (*viper.Viper)(p)
}

func (p *Postgres) Host() string {
	return p.viper().GetString("api.db.postgres.host")
}

func (p *Postgres) SetHost(host string) {
	p.viper().Set("api.db.postgres.host", host)
}

func (p *Postgres) Port() string {
	return p.viper().GetString("api.db.postgres.port")
}

func (p *Postgres) SetPort(port string) {
	p.viper().Set("api.db.postgres.port", port)
}

func (p *Postgres) Username() string {
	return p.viper().GetString("api.db.postgres.username")
}

func (p *Postgres) SetUsername(username string) {
	p.viper().Set("api.db.postgres.username", username)
}

func (p *Postgres) Password() string {
	return p.viper().GetString("api.db.postgres.password")
}

func (p *Postgres) SetPassword(password string) {
	p.viper().Set("api.db.postgres.password", password)
}

func (p *Postgres) DBName() string {
	return p.viper().GetString("api.db.postgres.dbName")
}

func (p *Postgres) SetDBName(password string) {
	p.viper().Set("api.db.postgres.dbName", password)
}
