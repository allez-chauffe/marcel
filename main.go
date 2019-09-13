package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/Zenika/marcel/cmd"
)

func main() {
	if err := cmd.Marcel.Execute(); err != nil {
		log.Fatal(err)
	}
}
