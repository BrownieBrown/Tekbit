package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"server/data_analysis/helpers"
	"server/data_analysis/models"
	"server/logger"
)

func FetchStockData(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Log.Error().
			Str("url", url).
			Msgf("Failed to fetch stock data: %v", err)
		return nil, err
	}
	return resp, nil
}

func GetStockData(c echo.Context) error {
	symbol := c.Param("symbol")
	apiKey := os.Getenv("ALPHA_VANTAGE_API_KEY")
	params := map[string]string{
		"function": "TIME_SERIES_INTRADAY",
		"symbol":   symbol,
		"interval": "5min",
		"apikey":   apiKey,
	}
	url := helpers.ConstructAlphaVantageURL(params)

	logger.Log.Info().
		Str("symbol", symbol).
		Msg("Fetching stock data")

	// HTTP request
	resp, err := FetchStockData(url)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to fetch stock data: %v", err))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Error().
			Msgf("Failed to read response body: %v", err)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to read response body: %v", err))
	}

	var stockData models.Stock
	err = json.Unmarshal(body, &stockData)
	if err != nil {
		logger.Log.Error().
			Msgf("Failed to unmarshal response body: %v", err)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to unmarshal response body: %v", err))
	}

	return c.JSON(http.StatusOK, stockData)
}
