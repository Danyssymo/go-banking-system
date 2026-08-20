package main

import (
	"fmt"
	"log"

	"github.com/Danyssymo/go-banking-system/user-service/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	fmt.Printf("user-service: starting in %s mode on port %d\n", cfg.Environment, cfg.HTTPPort)
}
