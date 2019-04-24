package main

import (
	"log"

	"github.com/Zenika/MARCEL/backend/cmd"
)

func main() {
	if err := cmd.Cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
