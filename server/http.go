package server

import (
	"github.com/mosescode1/Proxime/internal/middleware"
	"github.com/mosescode1/Proxime/internal/routes"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/mosescode1/Proxime/config"
)

func StartServer(config *config.Config) {
	app := echo.New()

	// Middlewares
	middleware.SetUpMiddleware(app)
	routes.SetUpRoutes(app)
	log.Fatal(app.Start(":" + config.APP_PORT)) // Listen on port 8080
}
