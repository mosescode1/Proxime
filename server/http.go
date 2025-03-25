package server

import (
	"log"

	"github.com/labstack/echo"
	"github.com/mosescode1/Proxime/config"
)

func StartServer(config *config.Config) {
	app := echo.New()

	// app.GET("/")

	log.Fatal(app.Start(":" + config.APP_PORT)) // Listen on port 8080
}
