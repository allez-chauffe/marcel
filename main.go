package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/Zenika/MARCEL/cmd"
)

func main() {
	if err := cmd.Marcel.Execute(); err != nil {
		log.Fatal(err)
	}
}
