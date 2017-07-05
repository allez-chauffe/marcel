package main

//go:generate swagger generate spec

import "github.com/Zenika/MARCEL/backend/app"

//todo : service to return a media configuration for the frontend (ie: everything but the back props for plugins)
//todo : service to create a new media
//todo : service to log with jwt (or at least a token based system)
//todo : service to delete a media
//todo : service to save an existing media
//todo : service to return list of all plugins
//todo : service to add a plugin to a media
//todo : plugin : import a new one
//todo : plugin : run plugins

func main() {

	a := new(app.App)
	a.Initialize()

	a.Run(":8090")
}
