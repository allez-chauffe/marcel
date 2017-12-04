package main

import (
	"github.com/Zenika/MARCEL/auth-backend/app"
	"github.com/Zenika/MARCEL/auth-backend/conf"
	"github.com/Zenika/MARCEL/auth-backend/users"
)

func main() {
	config := conf.LoadConfig()
	users.LoadUsersData()
	app.Run(config)
}
