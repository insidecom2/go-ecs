package main

import (
	"encoding/json"
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

	g.GET("/symbols", func(c *echo.Context) error {
		symbols := getSymbols()
		return c.JSON(http.StatusOK, symbols)
	})

	if err := c.Start(":8080"); err != nil {
		slog.Error("server stopped", "error", err)
	}
}

func getSymbols() []string {
	url := "https://api.kraken.com/0/public/Assets"

	resp, err := http.Get(url)
	if err != nil {
		slog.Error("failed to fetch symbols", "error", err)
		return nil
	}
	defer resp.Body.Close()

	var result struct {
		Result map[string]interface{} `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		slog.Error("failed to decode response", "error", err)
		return nil
	}

	symbols := make([]string, 0, len(result.Result))
	for symbol := range result.Result {
		symbols = append(symbols, symbol)
	}
	return symbols
}
