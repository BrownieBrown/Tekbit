package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"server/data_analysis/api"
	"server/utils"
)

func main() {
	utils.LoadEnvironment()

	e := echo.New()

	api.SetRoutes(e)

	err := e.Start(":8000")
	if err != nil {
		log.Fatal("Unable to start the server. ", err)
	}
}
