package karken

import (
	"go-ecs/config"
	"go-ecs/internal/adaptors"
	"net/http"

	"github.com/labstack/echo/v5"
)

type KarkenHandle struct {
	config  *config.Config
	adaptor *adaptors.KarkenAdaptor
}

type KarkenResp struct {
	Symbol  string `json:"symbol"`
	Decimal int16  `json:"decimal"`
}

type Karken interface {
	GetSymbols(c *echo.Context) error
}

func New(c *config.Config, adaptor *adaptors.KarkenAdaptor) *KarkenHandle {
	return &KarkenHandle{
		config:  c,
		adaptor: adaptor,
	}
}

func (h *KarkenHandle) GetSymbols(c *echo.Context) error {
	result := h.adaptor.GetSymbols()
	var symbols []KarkenResp

	for symbol, value := range result {
		symbols = append(symbols, KarkenResp{
			Symbol:  symbol,
			Decimal: int16(value.Decimals),
		})
	}

	if symbols == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Notfound"})
	}
	return c.JSON(http.StatusOK, symbols)
}
