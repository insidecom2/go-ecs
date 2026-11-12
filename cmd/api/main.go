package main

import (
	"go-ecs/config"
	"go-ecs/internal/models"
	"go-ecs/internal/module/health"
	"go-ecs/internal/module/karken"
	"log/slog"
	"os"

	"github.com/golang-cz/devslog"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	c := echo.New()
	slog.SetDefault(slog.New(devslog.NewHandler(os.Stdout, nil)))
	c.Logger = slog.Default()

	config := loadEnvConfig()
	db, err := dbConnect(config)
	if err != nil {
		slog.Error("database connection failed", "error", err)
		return
	}

	if err := db.AutoMigrate(&models.Symbols{}); err != nil {
		slog.Error("database migration failed", "error", err)
		return
	}

	g := c.Group("/api/v1")
	health.Route(g)
	karken.Route(g, config, db)

	if err := c.Start(":8080"); err != nil {
		slog.Error("server stopped", "error", err)
	}
}

func dbConnect(config *config.Config) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(config.PostgresUrl), &gorm.Config{})
}

func loadEnvConfig() *config.Config {
	if err := godotenv.Load(); err != nil {
		slog.Warn("No .env file found; using system environment variables")
	}
	return &config.Config{
		KarkenUrl:   os.Getenv("KARKEN_URL"),
		PostgresUrl: os.Getenv("POSTGRES_URL"),
	}
}
