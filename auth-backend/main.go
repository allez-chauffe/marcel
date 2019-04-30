package main

import (
	"log"

	"github.com/Zenika/MARCEL/auth-backend/cmd"
)

func main() {
	if err := cmd.AuthCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
