package karken

import (
	"go-ecs/config"
	"go-ecs/internal/adaptors"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func Route(g *echo.Group, config *config.Config, db *gorm.DB) {
	adaptor := adaptors.NewKarkenAdaptor(config)
	postgres := adaptors.NewPostgres(db)
	handle := New(config, adaptor, postgres)

	g.GET("/symbols", handle.GetSymbols)
}
