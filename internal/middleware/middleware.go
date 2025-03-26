package middleware

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetUpMiddleware(app *echo.Echo) {

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	app.Use(middleware.CSRF())
	app.Use(middleware.RequestID())
	app.Use(middleware.Secure())
	app.Use(echoprometheus.NewMiddleware("chat_app"))

}
