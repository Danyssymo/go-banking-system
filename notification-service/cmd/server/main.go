package main

import (
	"fmt"
	"log"

	"github.com/Danyssymo/go-banking-system/notification-service/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	fmt.Printf("notification-service: starting in %s mode on port %d\n", cfg.Environment, cfg.HTTPPort)
}
