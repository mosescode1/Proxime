package routes

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/mosescode1/Proxime/internal/handler"
	"github.com/mosescode1/Proxime/internal/socket"
)

func SetUpRoutes(app *echo.Echo) {
	chatHandler := handler.NewChatHandler()
	app.GET("/metrics", echoprometheus.NewHandler())
	app.GET("/health", chatHandler.Health)
}

func SetUpSocketRoutes(app *echo.Echo) {
	socketHandler := socket.NewSocket()
	app.GET("/chat", socketHandler.Connect)
}
