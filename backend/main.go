//go:generate swagger generate spec -o ./swagger.json

// Package main MARCEL APIs
//
// Provide API to access medias information
//
//     Host: localhost:8090
//     BasePath: /api/v1/
//     Version: 1.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Gwennael Buchet<gwennael.buchet@gmail.com>
//
// swagger:meta
package main

import "github.com/Zenika/MARCEL/backend/app"

//todo : service to log with jwt (or at least a token based system)
//todo : service to delete a media
//todo : service to add a plugin to a media
//todo : plugin : import/upload a new one
//todo : plugin : run plugins
//todo : service to upload a file for a specific plugin instance while saving the configuration

func main() {

	a := new(app.App)
	a.Initialize()

	a.Run(":8090")
}
