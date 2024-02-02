package main

import (
	"discord-go-bot/config"
	"discord-go-bot/internal/applicator"
	"log"
)

// Project entry point
func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	applicator.Run(cfg)

}
