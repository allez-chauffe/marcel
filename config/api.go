package config

import (
	"os"
	"path/filepath"
	"time"
)

type privateAPI struct {
	Port       uint
	BasePath   string
	CORS       bool
	DBFile     string
	PluginsDir string
	MediasDir  string
	DataDir    string
	Auth       privateAuth
}

type api struct {
	*privateAPI
	Auth auth
}

type privateAuth struct {
	Secure            bool
	Expiration        time.Duration
	RefreshExpiration time.Duration
}

type auth struct {
	*privateAuth
}

func (apiConfig api) getPath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return os.ExpandEnv(filepath.Join(apiConfig.privateAPI.DataDir, path))
}

func (apiConfig api) PluginsDir() string {
	return apiConfig.getPath(apiConfig.privateAPI.PluginsDir)
}

func (apiConfig api) MediasDir() string {
	return apiConfig.getPath(apiConfig.privateAPI.MediasDir)
}

func (apiConfig api) DataDir() string {
	return os.ExpandEnv(apiConfig.privateAPI.DataDir)
}

func (apiConfig api) DBFile() string {
	return apiConfig.getPath(apiConfig.privateAPI.DBFile)
}
