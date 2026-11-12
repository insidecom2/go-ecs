package karken

import (
	"go-ecs/config"
	"go-ecs/internal/adaptors"

	"github.com/labstack/echo/v5"
)

func Route(g *echo.Group, config *config.Config) {
	adaptor := adaptors.New(config)
	handle := New(config, adaptor)

	g.GET("/symbols", handle.GetSymbols)
}
