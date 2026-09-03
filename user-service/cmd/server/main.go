package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Danyssymo/go-banking-system/user-service/internal/adapters/hasher"
	httpadapter "github.com/Danyssymo/go-banking-system/user-service/internal/adapters/http"
	"github.com/Danyssymo/go-banking-system/user-service/internal/adapters/postgres"
	"github.com/Danyssymo/go-banking-system/user-service/internal/config"
	"github.com/Danyssymo/go-banking-system/user-service/internal/usecase"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to create connection pool: %v", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	userRepo := postgres.NewUserRepository(pool)
	passwordHasher := hasher.NewBcryptHasher()

	registerUC := usecase.NewRegisterUseCase(userRepo, passwordHasher)
	registerHandler := httpadapter.NewRegisterHandler(registerUC)

	router := gin.Default()
	router.POST("/register", registerHandler.Handle)

	log.Printf("user-service: db connected, environment=%s, port=%d", cfg.Environment, cfg.HTTPPort)

	addr := fmt.Sprintf(":%d", cfg.HTTPPort)
	if err := router.Run(addr); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
