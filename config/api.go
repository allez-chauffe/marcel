package config

import (
	"os"
	"path/filepath"
	"time"
)

type api struct {
	Port       uint
	BasePath   string
	CORS       bool
	DBFile     string
	PluginsDir string
	MediasDir  string
	DataDir    string
	Auth       auth
}

type auth struct {
	Secure            bool
	Expiration        time.Duration
	RefreshExpiration time.Duration
}

func (apiConfig api) getPath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return os.ExpandEnv(filepath.Join(apiConfig.DataDir, path))
}

func (apiConfig api) GetPluginsDir() string {
	return apiConfig.getPath(apiConfig.PluginsDir)
}

func (apiConfig api) GetMediasDir() string {
	return apiConfig.getPath(apiConfig.MediasDir)
}

func (apiConfig api) GetDataDir() string {
	return os.ExpandEnv(apiConfig.DataDir)
}

func (apiConfig api) GetDBFile() string {
	return apiConfig.getPath(apiConfig.DBFile)
}
