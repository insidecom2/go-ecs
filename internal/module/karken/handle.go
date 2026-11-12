package karken

import (
	"go-ecs/config"
	"go-ecs/internal/adaptors"
	"go-ecs/internal/models"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v5"
)

type KarkenHandle struct {
	config   *config.Config
	adaptor  *adaptors.KarkenAdaptor
	postgres *adaptors.Postgres
}

type KarkenResp struct {
	Symbol  string `json:"symbol"`
	Decimal int16  `json:"decimal"`
}

type Karken interface {
	GetSymbols(c *echo.Context) error
}

func New(c *config.Config, adaptor *adaptors.KarkenAdaptor, postgres *adaptors.Postgres) *KarkenHandle {
	return &KarkenHandle{
		config:   c,
		adaptor:  adaptor,
		postgres: postgres,
	}
}

func (h *KarkenHandle) GetSymbols(c *echo.Context) error {
	result := h.adaptor.GetSymbols()
	var symbols []KarkenResp
	var records []models.Symbols

	for symbol, value := range result {
		records = append(records, models.Symbols{
			Symbol:            symbol,
			Name:              symbol,
			CirculatingSupply: 0.00,
			TotalSupply:       0.00,
			MaxSupply:         0.00,
			LastPrice:         0.00,
		})
		symbols = append(symbols, KarkenResp{
			Symbol:  symbol,
			Decimal: int16(value.Decimals),
		})
	}
	if err := h.postgres.UpsertBatch(c.Request().Context(), records); err != nil {
		slog.Error("failed to save symbols", "error", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to save symbols",
		})
	}

	if symbols == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Notfound"})
	}
	return c.JSON(http.StatusOK, symbols)
}
