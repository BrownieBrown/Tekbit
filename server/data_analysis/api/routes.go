package api

import (
	"github.com/labstack/echo/v4"
	"server/data_analysis/api/handlers"
)

func SetRoutes(e *echo.Echo) {
	e.GET("/api/v1/stocks/:symbol", handlers.GetStockData)
}
