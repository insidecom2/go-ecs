package adaptors

import (
	"encoding/json"
	"fmt"
	"go-ecs/config"
	"log/slog"
	"net/http"
)

type KarkenAdaptor struct {
	config *config.Config
}

type AssetResponse struct {
	Errors []string
	Result map[string]Asset
}

type Asset struct {
	AClass          string `json:"aclass"`
	AltName         string `json:"altname"`
	Decimals        int    `json:"decimals"`
	DisplayDecimals int    `json:"display_decimals"`
	Status          string `json:"status"`
	MarginRate      string `json:"margin_rate"`
}

type Karken interface {
	GetSymbols() map[string]Asset
}

func NewKarkenAdaptor(c *config.Config) *KarkenAdaptor {
	return &KarkenAdaptor{
		config: c,
	}
}

func (a KarkenAdaptor) GetSymbols() map[string]Asset {
	url := fmt.Sprintf("%s/0/public/Assets", a.config.KarkenUrl)

	resp, err := http.Get(url)
	if err != nil {
		slog.Error("failed to fetch symbols", "error", err)
		return nil
	}
	defer resp.Body.Close()

	var result AssetResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		slog.Error("failed to decode response", "error", err)
		return nil
	}

	return result.Result
}
