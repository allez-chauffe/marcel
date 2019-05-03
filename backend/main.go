package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/Zenika/MARCEL/backend/cmd"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	if err := cmd.Marcel.Execute(); err != nil {
		log.Fatal(err)
	}
}
