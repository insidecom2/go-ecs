package health

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func Route(g *echo.Group) {
	g.GET("/health", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "OK",
		})
	})
}
