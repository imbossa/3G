package main

import (
	"log"

	"github.com/imbossa/3G/config"
	"github.com/imbossa/3G/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
