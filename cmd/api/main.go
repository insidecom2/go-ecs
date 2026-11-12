package main

import (
	"go-ecs/config"
	"go-ecs/internal/module/health"
	"go-ecs/internal/module/karken"
	"log/slog"
	"os"

	"github.com/golang-cz/devslog"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
)

func main() {
	c := echo.New()
	slog.SetDefault(slog.New(devslog.NewHandler(os.Stdout, nil)))
	c.Logger = slog.Default()

	config := loadEnvConfig()

	g := c.Group("/api/v1")
	health.Route(g)
	karken.Route(g, config)

	if err := c.Start(":8080"); err != nil {
		slog.Error("server stopped", "error", err)
	}
}

func loadEnvConfig() *config.Config {
	if err := godotenv.Load(); err != nil {
		slog.Warn("No .env file found; using system environment variables")
	}
	return &config.Config{
		KarkenUrl: os.Getenv("KARKEN_URL"),
	}
}
