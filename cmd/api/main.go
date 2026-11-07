package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/golang-cz/devslog"
	"github.com/labstack/echo/v5"
)

func main() {
	c := echo.New()
	slog.SetDefault(slog.New(devslog.NewHandler(os.Stdout, nil)))
	c.Logger = slog.Default()

	g := c.Group("/api/v1")
	g.GET("/health", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"messsage": "OK",
		})
	})

	if err := c.Start(":8080"); err != nil {
		slog.Error("server stopped", "error", err)
	}
}
