package main

import (
	"github.com/Zenika/MARCEL/auth-backend/app"
	"github.com/Zenika/MARCEL/auth-backend/conf"
)

func main() {
	config := conf.LoadConfig()
	app.Run(config)
}
